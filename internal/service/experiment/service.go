package experiment

type Service struct {
	experimentRepository experimentRepository
}

func NewService(experimentRepository experimentRepository) *Service {
	return &Service{
		experimentRepository: experimentRepository,
	}
}
