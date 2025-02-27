package entities

import (
	"transactions/domain/vo"
)

// O balanço absoluto da conta é resultante da compensação de todas as transações

type Account struct {
	id            vo.AccountID
	ownerDocument vo.AccountOwnerDocument
}

func NewAccount(ownerDocument vo.AccountOwnerDocument) Account {
	return Account{
		id:            vo.NewAccountID(),
		ownerDocument: ownerDocument,
	}
}

func (a *Account) OwnerDocument() vo.AccountOwnerDocument {
	return a.ownerDocument
}

func (a *Account) CalculateBalance(transactions []Transaction) int {
	balance := 0
	for _, transaction := range transactions {
		if transaction.transactionType == vo.TransactionTypeCredit {
			balance += transaction.amount
		} else if transaction.transactionType == vo.TransactionTypeDebit {
			balance -= transaction.amount
		}
	}
	return balance
}
