package message

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gin-gorm-sample/application/domain"
	messagePort "gin-gorm-sample/mocks/application/message/port"
	userPort "gin-gorm-sample/mocks/application/user/port"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessagePort := messagePort.NewMockRepository(ctrl)
	mockUserPort := userPort.NewMockRepository(ctrl)

	service := NewService(mockMessagePort, mockUserPort)

	t.Run("should be error if user id check error", func(t *testing.T) {
		mockUserPort.EXPECT().Get(gomock.Any()).Return(domain.User{}, errors.New("repository error"))
		message := domain.Message{}
		response := service.Create(&message)

		assert.Equal(t, errors.New("repository error"), response)
	})

	t.Run("should be pass", func(t *testing.T) {
		mockUserPort.EXPECT().Get(gomock.Any()).Return(domain.User{}, nil)
		mockMessagePort.EXPECT().Create(gomock.Any()).Return(nil)

		message := domain.Message{}
		response := service.Create(&message)

		assert.Equal(t, nil, response)
	})
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMessagePort := messagePort.NewMockRepository(ctrl)
	mockUserPort := userPort.NewMockRepository(ctrl)

	service := NewService(mockMessagePort, mockUserPort)

	t.Run("should be error if user id check error", func(t *testing.T) {
		mockUserPort.EXPECT().Get(gomock.Any()).Return(domain.User{}, errors.New("repository error"))
		message := domain.Message{UserID: 123}
		response := service.Update(&message)

		assert.Equal(t, errors.New("repository error"), response)
	})

	t.Run("should be pass", func(t *testing.T) {
		mockUserPort.EXPECT().Get(gomock.Any()).Return(domain.User{}, nil)
		mockMessagePort.EXPECT().Create(gomock.Any()).Return(nil)

		message := domain.Message{UserID: 123}
		response := service.Create(&message)

		assert.Equal(t, nil, response)
	})
}
