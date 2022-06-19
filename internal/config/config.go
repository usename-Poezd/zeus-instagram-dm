package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ZeusConfig
}

type ZeusConfig struct {
	Token   string `mapstructure:"ZEUS_TOKEN"`
}

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init(configFile string) (*Config, error) {

	if err := parseConfigFile(configFile); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}


	return &cfg, nil
}

func parseConfigFile(configFile string) error {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	
	if err := viper.ReadInConfig(); err != nil {
		return err
	}


	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	return viper.UnmarshalExact(&cfg.ZeusConfig)
}