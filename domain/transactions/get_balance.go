package transactions

import (
	"transactions/domain/vo"
)

type GetBalanceUC struct {
	transactionRepository transactionRepository
	accountRepository     accountRepository
}

type GetBalanceInput struct {
	accountOwnerDocument string
}

func (g GetBalanceUC) Get(input GetBalanceInput) (int, error) {
	accountOwnerDocument, err := vo.ParseAccountOwnerDocument(input.accountOwnerDocument)
	if err != nil {
		return 0, err
	}

	account, err := g.accountRepository.GetAccountByDocument(accountOwnerDocument)
	if err != nil {
		return 0, err
	}

	transactions, err := g.transactionRepository.GetTransactionsByDocument(accountOwnerDocument)
	if err != nil {
		return 0, err
	}

	return account.CalculateBalance(transactions), nil

}
