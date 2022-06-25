package services

type ZeusService struct {
	Token  string
	ApiUrl string
}

func NewZeusService(Token string, ApiUrl string) *ZeusService {
	return &ZeusService{
		Token,
		ApiUrl,
	}
}

func (s *ZeusService) GetChampagin(id int) error {
	return nil
}

func (s *ZeusService) CreateChampagin() error {

	return nil
}

func (s *ZeusService) UpdateChampagin(id int) error {
	return nil
}

func (s *ZeusService) DeleteChampagin(id int) error {
	return nil
}

func (s *ZeusService) StartChampagin(id int) error {
	return nil
}

func (s *ZeusService) StopChampagin(id int) error {
	return nil
}
