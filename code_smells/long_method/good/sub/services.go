package sub

import "errors"

type services struct {
	repo *repositories
}

func NewServices(repo *repositories) *services {
	return &services{
		repo: repo,
	}
}

func (s *services) GetUser(userID int) (responseUser, error) {
	response, err := s.repo.GetUser(userID)
	if err != nil {
		return responseUser{}, err
	}

	if !response.Status.isActive() {
		return responseUser{}, errors.New("user not active")
	}

	return response.userToResponseUser(), nil
}
