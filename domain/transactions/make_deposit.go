package transactions

import (
	"fmt"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type MakeDepositUC struct {
	transactionRepository transactionRepository
	accountRepository     accountRepository
}

type ExecuteDepositInput struct {
	UserAccountOwnerDocument string
	Amount                   int
}

func (m MakeDepositUC) ExecuteDeposit(input ExecuteDepositInput) error {
	ownerDocument, err := vo.ParseAccountOwnerDocument(input.UserAccountOwnerDocument)
	if err != nil {
		return err
	}

	_, err = m.accountRepository.GetAccountByDocument(ownerDocument)
	if err != nil {
		return fmt.Errorf("the account has no balance")
	}

	transaction := entities.NewTransaction(vo.RelatedTransactionID{}, ownerDocument, vo.TransactionTypeCredit, input.amount)
	if err := m.transactionRepository.InsertTransaction(transaction); err != nil {
		return err
	}

	return nil
}
