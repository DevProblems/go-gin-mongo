package services

import "github.com/DevProblems/sarang-gin-mongo-apis/models"

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
