package user

type Service struct {
	userRepository userRepository
}

func NewService(userRepository userRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
