package transactions

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

func TestMakeDepositUC_ExecuteDeposit(t *testing.T) {
	t.Run("Should execute a deposit", func(t *testing.T) {
		mockTransactionRepo := new(MockTransactionRepository)
		mockAccountRepo := new(MockAccountRepository)

		depositUC := MakeDepositUC{
			transactionRepository: mockTransactionRepo,
			accountRepository: mockAccountRepo,
		}

		input := ExecuteDepositInput{
			accountOwnerDocument: "12345678901",
			amount:               1000,
		}

		// Account
		accountOwnerDocument, err := vo.ParseAccountOwnerDocument(input.accountOwnerDocument)
		require.NoError(t, err)
		account := entities.NewAccount(accountOwnerDocument)

		// Transaction
		creditTransaction := entities.NewTransaction(vo.RelatedTransactionID{}, account.OwnerDocument(), vo.TransactionTypeCredit, input.amount)

		mockTransactionRepo.On("InsertTransaction", mock.Anything).Return(creditTransaction, nil)
		mockAccountRepo.On("GetAccountByDocument", mock.Anything).Return(account, nil)

		err = depositUC.ExecuteDeposit(input)
		require.NoError(t, err)
		mockAccountRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})
}
