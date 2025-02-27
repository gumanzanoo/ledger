package transactions

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

func TestGetTransactionsUC_GetTransactions(t *testing.T) {
	t.Run("Should get transactions", func(t *testing.T) {
		mockTransactionRepo := new(MockTransactionRepository)
		mockAccountRepo := new(MockAccountRepository)

		getTransactionsUC := NewGetTransactionsUC(mockTransactionRepo, mockAccountRepo)

		input := GetTransactionsInput{accountOwnerDocument: "12345678901"}

		// Account
		accountOwnerDocument, err := vo.ParseAccountOwnerDocument(input.accountOwnerDocument)
		require.NoError(t, err)

		account := entities.NewAccount(accountOwnerDocument)

		// Transactions
		transactions := []entities.Transaction{
			entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, 1000),
			entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, 3000),
			entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, 5000),
		}

		mockAccountRepo.On("GetOwnerByDocument", mock.Anything).Return(account, nil)
		mockTransactionRepo.On("GetTransactionsByDocument", mock.Anything).Return(transactions, nil)

		_, err = getTransactionsUC.GetTransactions(input)
		require.NoError(t, err)

		mockAccountRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})
}
