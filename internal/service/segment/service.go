package segment

type Service struct {
	segmentRepository segmentRepository
}

func NewService(segmentRepository segmentRepository) *Service {
	return &Service{
		segmentRepository: segmentRepository,
	}
}
