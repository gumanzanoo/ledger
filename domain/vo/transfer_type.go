package vo

import (
	"fmt"
	"strings"
)

var (
	ErrorInvalidTransactionType = fmt.Errorf("invalid transaction type")
)

type TransactionType struct {
	value string
}

var (
	TransactionTypeDebit  = TransactionType{"debit"}
	TransactionTypeCredit = TransactionType{"credit"}
)

func (t TransactionType) Value() string {
	return t.value
}

func (t TransactionType) IsDebit() bool {
	return t.value == TransactionTypeDebit.value
}

func (t TransactionType) IsCredit() bool {
	return t.value == TransactionTypeCredit.value
}

func ParseTransactionType(s string) (TransactionType, error) {
	switch strings.ToLower(s) {
	case TransactionTypeDebit.value:
		return TransactionTypeDebit, nil
	case TransactionTypeCredit.value:
		return TransactionTypeCredit, nil
	default:
		return TransactionType{}, ErrorInvalidTransactionType
	}
}
