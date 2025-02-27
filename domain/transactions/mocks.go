package transactions

import (
	"github.com/stretchr/testify/mock"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) InsertTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionsByDocument(document vo.AccountOwnerDocument) ([]entities.Transaction, error) {
	args := m.Called(document)
	return args.Get(0).([]entities.Transaction), args.Error(1)
}

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) GetOwnerByDocument(document vo.AccountOwnerDocument) (entities.Account, error) {
	args := m.Called(document)
	return args.Get(0).(entities.Account), args.Error(1)
}
