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

func NewMakeDepositUC(
	transactionRepository transactionRepository,
	accountRepository accountRepository,
) MakeDepositUC {
	return MakeDepositUC{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

type ExecuteDepositInput struct {
	accountOwnerDocument string
	amount               int
}

func (m MakeDepositUC) ExecuteDeposit(input ExecuteDepositInput) error {
	ownerDocument, err := vo.ParseAccountOwnerDocument(input.accountOwnerDocument)
	if err != nil {
		return err
	}

	_, err = m.accountRepository.GetOwnerByDocument(ownerDocument)
	if err != nil {
		return fmt.Errorf("the account has no balance")
	}

	transaction := entities.NewTransaction(vo.RelatedTransactionID{}, ownerDocument, vo.TransactionTypeCredit, input.amount)
	_, err = m.transactionRepository.InsertTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}
