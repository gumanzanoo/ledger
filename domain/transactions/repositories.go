package transactions

import (
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type transactionRepository interface {
	InsertTransaction(transaction entities.Transaction) (entities.Transaction, error)
	GetTransactionsByDocument(document vo.AccountOwnerDocument) ([]entities.Transaction, error)
}

type accountRepository interface {
	GetOwnerByDocument(document vo.AccountOwnerDocument) (entities.Account, error)
}
