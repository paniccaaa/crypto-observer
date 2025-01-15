package v1

type Storage interface {
}

type Service struct {
	cryptoRepository Storage
}

func NewService(cryptoRepo Storage) *Service {
	return &Service{
		cryptoRepository: cryptoRepo,
	}
}
