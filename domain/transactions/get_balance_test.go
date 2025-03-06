package transactions

import (
	"testing"
	"transactions/domain/entities"
	"transactions/domain/vo"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetBalanceUC_GetBalance(t *testing.T) {
	mockTransactionRepository := new(MockTransactionRepository)
	mockAccountRepository := new(MockAccountRepository)

	getBalanceUC := GetBalanceUC{
		transactionRepository: mockTransactionRepository,
		accountRepository: mockAccountRepository,
	}

	input := GetBalanceInput{accountOwnerDocument: "12345678901"}

	// account mock
	accountOwnerDocument, err := vo.ParseAccountOwnerDocument(input.accountOwnerDocument)
	require.NoError(t, err)

	account := entities.NewAccount(accountOwnerDocument)

	// transactions mock
	transactions := []entities.Transaction{
		entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, 1000),
		entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, 3000),
		entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, 5000),
	}

	mockAccountRepository.On("GetAccountByDocument", mock.Anything).Return(account, nil)
	mockTransactionRepository.On("GetTransactionsByDocument", mock.Anything).Return(transactions, nil)

	_, err = getBalanceUC.Get(input)
	require.NoError(t, err)

	mockAccountRepository.AssertExpectations(t)
	mockAccountRepository.AssertExpectations(t)
}
