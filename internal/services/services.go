package services

type Deps struct {
    ZeusApiUrl string
    ZeusToken string
}

type Services struct {
    Zeus Zeus
}

type Zeus interface {
    GetChampagin(id int) error
    CreateChampagin() error
    UpdateChampagin(id int) error
    DeleteChampagin(id int) error
    
    StartChampagin(id int) error
    StopChampagin(id int) error
}

func NewServices(deps Deps) *Services {
	return &Services{
		Zeus: NewZeusService(deps.ZeusToken, deps.ZeusApiUrl),
	}
}