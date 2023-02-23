package queueRabbitProcessRepository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/workerRabbitMQ-example/adapter/postgres/queueRabbitProcessRepository"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupCreate() ([]string, dto.QueueProcessDTO, domain.QueueProcess, pgxmock.PgxPoolIface) {
	cols := []string{"id", "queuemessage", "message", "result", "created_at"}
	fakeQueueProcessDTO := dto.QueueProcessDTO{}
	fakeQueueProcess := domain.QueueProcess{}
	faker.FakeData(&fakeQueueProcessDTO)
	faker.FakeData(&fakeQueueProcess)

	mock, _ := pgxmock.NewPool()

	return cols, fakeQueueProcessDTO, fakeQueueProcess, mock
}

func TestCreate(t *testing.T) {
	cols, fakeQueueProcessDTO, fakeQueueProcess, mock := setupCreate()
	defer mock.Close()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO queue_message_process (queuemessage,message,result) VALUES ($1, $2, $3) returning *")).WithArgs(
		fakeQueueProcessDTO.QueueMessage,
		fakeQueueProcessDTO.Message,
		fakeQueueProcessDTO.Result,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeQueueProcess.ID,
		fakeQueueProcess.QueueMessage,
		fakeQueueProcess.Message,
		fakeQueueProcess.Result,
		fakeQueueProcess.CreatedAt,
	))

	sut := queueRabbitProcessRepository.New(mock)
	queueRet, err := sut.Create(&fakeQueueProcessDTO)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, queueRet.ID)
	require.Equal(t, queueRet.Message, fakeQueueProcess.Message)
	require.Equal(t, queueRet.Result, fakeQueueProcess.Result)

}

func TestCreate_DBError(t *testing.T) {
	_, fakeQueueProcessDTO, _, mock := setupCreate()
	fakeQueueProcessDTO.Result = "T"
	defer mock.Close()

	mock.ExpectQuery("INSERT INTO queue_message_process (queuemessage,message,result) VALUES ($1, $2, $3)").WithArgs(
		fakeQueueProcessDTO.QueueMessage,
		fakeQueueProcessDTO.Message,
		fakeQueueProcessDTO.Result,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := queueRabbitProcessRepository.New(mock)
	queueRet, err := sut.Create(&fakeQueueProcessDTO)

	require.NotNil(t, err)
	require.Nil(t, queueRet)
}
