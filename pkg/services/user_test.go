package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xion6/golang-echo/pkg/config"
	"github.com/xion6/golang-echo/pkg/domain"
	usermocks "github.com/xion6/golang-echo/pkg/mocks/data/users"
)

func TestCreateAccount_UserExists(t *testing.T) {
	cfg := &config.Settings{}
	userProviderMock := &usermocks.UserProviderMock{}
	userProviderMock.UsernameExistsMock = func(username string) (bool, error) {
		return true, nil
	}

	userSvc := NewUserService(cfg, userProviderMock)
	newUser := &domain.User{Username: "test", Password: "test"}

	response := userSvc.CreateAccount(newUser)
	assert.Equal(t, "USERNAME_TAKEN", response.Name)

}
