package user

import "example.com/users/internal/ports"

type Option func(*Service)
type Service struct {
	UserRepository ports.UserRepository
}

func NewService(opts ...Option) ports.UserService {
	s := &Service{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithRepository(repository ports.UserRepository) Option {
	return func(s *Service) {
		s.UserRepository = repository
	}
}
