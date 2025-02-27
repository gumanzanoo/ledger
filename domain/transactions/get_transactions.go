package transactions

import (
	"fmt"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type GetTransactionsUC struct {
	transactionRepository transactionRepository
	accountRepository     accountRepository
}

func NewGetTransactionsUC(
	transactionRepository transactionRepository,
	accountRepository accountRepository,
) GetTransactionsUC {
	return GetTransactionsUC{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

type GetTransactionsInput struct {
	accountOwnerDocument string
}

func (m GetTransactionsUC) GetTransactions(input GetTransactionsInput) ([]entities.Transaction, error) {
	ownerDocument, err := vo.ParseAccountOwnerDocument(input.accountOwnerDocument)
	if err != nil {
		return []entities.Transaction{}, err
	}

	_, err = m.accountRepository.GetOwnerByDocument(ownerDocument)
	if err != nil {
		return []entities.Transaction{}, fmt.Errorf("the account has no balance")
	}

	transactions, err := m.transactionRepository.GetTransactionsByDocument(ownerDocument)
	if err != nil {
		return []entities.Transaction{}, err
	}

	return transactions, nil
}
