package transactions

import (
	"fmt"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type MakeTransactionUC struct {
	transactionRepository transactionRepository
	accountRepository     accountRepository
}

func NewMakeTransactionUC(
	transactionRepository transactionRepository,
	accountRepository accountRepository,
) MakeTransactionUC {
	return MakeTransactionUC{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

type ExecuteTransactionInput struct {
	userOriginDocument      string
	userDestinationDocument string
	amount                  int
}

func (m MakeTransactionUC) ExecuteTransaction(input ExecuteTransactionInput) error {
	userOriginDocument, err := vo.ParseAccountOwnerDocument(input.userOriginDocument)
	if err != nil {
		return err
	}

	userDestinationDocument, err := vo.ParseAccountOwnerDocument(input.userDestinationDocument)
	if err != nil {
		return err
	}

	userOriginAccount, err := m.accountRepository.GetOwnerByDocument(userOriginDocument)
	if err != nil {
		return err
	}

	userDestinationAccount, err := m.accountRepository.GetOwnerByDocument(userDestinationDocument)
	if err != nil {
		return err
	}

	originTransactions, err := m.transactionRepository.GetTransactionsByDocument(userOriginDocument)
	if err != nil {
		return err
	}

	if userOriginAccount.CalculateBalance(originTransactions) < input.amount {
		return fmt.Errorf("insufficient balance")
	}

	relatedTransactionID := vo.NewRelatedTransactionID()

	debitTransaction := entities.NewTransaction(relatedTransactionID, userOriginAccount.OwnerDocument(), vo.TransactionTypeDebit, input.amount)
	_, err = m.transactionRepository.InsertTransaction(debitTransaction)
	if err != nil {
		return err
	}

	creditTransaction := entities.NewTransaction(relatedTransactionID, userDestinationAccount.OwnerDocument(), vo.TransactionTypeCredit, input.amount)
	_, err = m.transactionRepository.InsertTransaction(creditTransaction)
	if err != nil {
		return err
	}

	return nil
}
