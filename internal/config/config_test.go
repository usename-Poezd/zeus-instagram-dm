package config

import (
	"os"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type env struct {
		Token string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("ZEUS_TOKEN", env.Token)
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				env: env{
					Token: "token",
				},
			},
			want: &Config{
				ZeusConfig{
					Token: "token",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)

			got, err := Init("../../.env")
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() Config error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() Config got = %v, want %v", got, tt.want)
			}
		})
	}
}