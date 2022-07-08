package rest

import (
	"WallE/models"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Register(user models.User) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockService) VerifikasiRegister(email, kode string) (string, error) {
	ret := m.Called()

	return ret.String(0), ret.Error(0)
}
func (m *MockService) GetUserDataById(id uint) (models.User, error) {
	ret := m.Called()

	return ret.Get(0).(models.User), ret.Error(0)
}
func (m *MockService) Login(email, password string) (string, int) {
	ret := m.Called()

	return ret.String(0), ret.Int(0)
}
func (m *MockService) CreateResetPassword(email string) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockService) UpdatePassword(email, password, code string) error {
	ret := m.Called()

	return ret.Error(0)
}
func (m *MockService) UpdateUserData(id uint, user models.User) error {
	ret := m.Called()

	return ret.Error(0)
}
