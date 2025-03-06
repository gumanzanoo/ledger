package transactions

import (
	"fmt"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type MakeTransactionUC struct {
	TransactionRepository transactionRepository
	AccountRepository     accountRepository
}

type ExecuteTransactionInput struct {
	UserOriginDocument      string
	UserDestinationDocument string
	Amount                  int
}

func (m MakeTransactionUC) ExecuteTransaction(input ExecuteTransactionInput) error {
	UserOriginDocument, err := vo.ParseAccountOwnerDocument(input.UserOriginDocument)
	if err != nil {
		return err
	}

	UserDestinationDocument, err := vo.ParseAccountOwnerDocument(input.UserDestinationDocument)
	if err != nil {
		return err
	}

	userOriginAccount, err := m.AccountRepository.GetAccountByDocument(UserOriginDocument)
	if err != nil {
		return err
	}

	userDestinationAccount, err := m.AccountRepository.GetAccountByDocument(UserDestinationDocument)
	if err != nil {
		return err
	}

	originTransactions, err := m.TransactionRepository.GetTransactionsByDocument(UserOriginDocument)
	if err != nil {
		return err
	}

	if userOriginAccount.CalculateBalance(originTransactions) < input.Amount {
		return fmt.Errorf("insufficient balance")
	}

	relatedTransactionID := vo.NewRelatedTransactionID()

	debitTransaction := entities.NewTransaction(relatedTransactionID, userOriginAccount.OwnerDocument(), vo.TransactionTypeDebit, input.Amount)
	if err := m.TransactionRepository.InsertTransaction(debitTransaction); err != nil {
		return err
	}

	creditTransaction := entities.NewTransaction(relatedTransactionID, userDestinationAccount.OwnerDocument(), vo.TransactionTypeCredit, input.Amount)
	if err := m.TransactionRepository.InsertTransaction(creditTransaction); err != nil {
		return err
	}

	return nil
}
