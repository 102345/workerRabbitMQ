package queueRabbitProcessUseCase_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/domain/mocks"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/marc/workerRabbitMQ-example/core/usecase/queueRabbitProcessUseCase"
	"github.com/stretchr/testify/require"
)

func TestCreateQueueProcess(t *testing.T) {
	fakeQueueProcessDTO := dto.QueueProcessDTO{}
	fakeQueueProcess := domain.QueueProcess{}
	faker.FakeData(&fakeQueueProcessDTO)
	faker.FakeData(&fakeQueueProcess)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQueueRabbitProcessRepository := mocks.NewMockIQueueProcessUseCase(mockCtrl)
	mockQueueRabbitProcessRepository.EXPECT().Create(&fakeQueueProcessDTO).Return(&fakeQueueProcess, nil)

	sut := queueRabbitProcessUseCase.New(mockQueueRabbitProcessRepository)
	queueRet, err := sut.Create(&fakeQueueProcessDTO)

	require.Nil(t, err)
	require.NotEmpty(t, queueRet.ID)
	require.Equal(t, queueRet.Message, fakeQueueProcess.Message)
	require.Equal(t, queueRet.Result, fakeQueueProcess.Result)
}

func TestCreateQueueProcess_Error(t *testing.T) {
	fakeQueueProcessDTO := dto.QueueProcessDTO{}
	faker.FakeData(&fakeQueueProcessDTO)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockQueueRabbitProcessRepository := mocks.NewMockIQueueProcessUseCase(mockCtrl)
	mockQueueRabbitProcessRepository.EXPECT().Create(&fakeQueueProcessDTO).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := queueRabbitProcessUseCase.New(mockQueueRabbitProcessRepository)
	queueRet, err := sut.Create(&fakeQueueProcessDTO)

	require.NotNil(t, err)
	require.Nil(t, queueRet)
}
