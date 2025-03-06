package entities

import (
	"fmt"
	"transactions/domain/vo"
)

// O balanço absoluto da conta é resultante da compensação de todas as transações

type Account struct {
	id            vo.AccountID
	ownerDocument vo.AccountOwnerDocument
}

func (a *Account) ID() vo.AccountID {
	return a.id
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

func (a *Account) ParseAccount(
	id string,
	ownerDocumentStr string,
) error {
	definitiveID, err := vo.ParseAccountID(id)
	if err != nil {
		return fmt.Errorf("could not parse account id %s", id)
	}

	definitiveOwnerDocument, err := vo.ParseAccountOwnerDocument(ownerDocumentStr)
	if err != nil {
		return fmt.Errorf("could not parse account owner document %s", ownerDocumentStr)
	}

	a.id = definitiveID
	a.ownerDocument = definitiveOwnerDocument
	return nil
}

func NewAccount(ownerDocument vo.AccountOwnerDocument) Account {
	return Account{
		id:            vo.NewAccountID(),
		ownerDocument: ownerDocument,
	}
}
