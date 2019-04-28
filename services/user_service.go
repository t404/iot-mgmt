package services

import (
	"iot-mgmt/models"
	"github.com/satori/go.uuid"
)

func NewCompanyService() *CompanyService {
	return &CompanyService{}
}

type CompanyService struct {
}

func (c *CompanyService) Create(request CreateUserRequest) (*models.User, error) {
	return &models.User{
		Id:   uuid.NewV4().String(),
		Name: request.Name,
	}, nil
}

func (c *CompanyService) Delete(request DeleteUserRequest) ([]models.User, error) {

	return func(request DeleteUserRequest) []models.User {
		var users []models.User
		for _, id := range request.Id {
			users = append(users, models.User{Id: id})
		}
		return users
	}(request), nil
}

type CreateUserRequest struct {
	// 用户名称
	Name string `json:"name" binding:"required"`
}

type DeleteUserRequest struct {
	// 要删除的用户 Id 列表
	Id []string `json:"id" binding:"required"`
}
