package transactions

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

func TestMakeOperationsUseCase_ExecuteTransaction(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockAccountRepo := new(MockAccountRepository)

	makeTransactionUC := NewMakeTransactionUC(mockTransactionRepo, mockAccountRepo)

	originAccountOwnerDocumentStr := "12345678901"
	destinationAccountOwnerDocumentStr := "10987654321"

	input := ExecuteTransactionInput{
		userOriginDocument:      originAccountOwnerDocumentStr,
		userDestinationDocument: destinationAccountOwnerDocumentStr,
		amount:                  10000,
	}

	// Origin account
	originAccountOwnerDocument, err := vo.ParseAccountOwnerDocument(input.userOriginDocument)
	require.NoError(t, err)
	originAccount := entities.NewAccount(originAccountOwnerDocument)
	oldTransactions := []entities.Transaction{
		entities.NewTransaction(
			vo.RelatedTransactionID{}, originAccount.OwnerDocument(), vo.TransactionTypeCredit, 50000),
	}

	// Destination account
	destinationAccountOwnerDocument, err := vo.ParseAccountOwnerDocument(input.userDestinationDocument)
	require.NoError(t, err)
	destinationAccount := entities.NewAccount(destinationAccountOwnerDocument)

	// Transactions
	relatedTransactionID := vo.NewRelatedTransactionID()
	debitTransaction := entities.NewTransaction(
		relatedTransactionID, originAccount.OwnerDocument(), vo.TransactionTypeDebit, input.amount)
	creditTransaction := entities.NewTransaction(
		relatedTransactionID, destinationAccount.OwnerDocument(), vo.TransactionTypeCredit, input.amount)

	mockAccountRepo.On("GetOwnerByDocument", mock.Anything).Return(originAccount, nil)
	mockAccountRepo.On("GetOwnerByDocument", mock.Anything).Return(destinationAccount, nil)
	mockTransactionRepo.On("GetTransactionsByDocument", mock.Anything).Return(oldTransactions, nil)
	mockTransactionRepo.On("InsertTransaction", mock.Anything).Return(debitTransaction, nil)
	mockTransactionRepo.On("InsertTransaction", mock.Anything).Return(creditTransaction, nil)

	err = makeTransactionUC.ExecuteTransaction(input)
	require.NoError(t, err)
	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}
