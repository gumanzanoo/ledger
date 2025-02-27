package vo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type transferTypeTest struct {
	transferTypeStr string
}

func TestParseTransactionType(t *testing.T) {
	validCreditStr := "credit"
	validDebitStr := "debit"

	t.Run("credit", func(t *testing.T) {
		t.Parallel()

		id, err := ParseTransactionType(validCreditStr)
		require.NoError(t, err)
		require.Equal(t, validCreditStr, id.Value())
		require.True(t, id.IsCredit())
	})

	t.Run("debit", func(t *testing.T) {
		t.Parallel()

		id, err := ParseTransactionType(validDebitStr)
		require.NoError(t, err)
		require.Equal(t, validDebitStr, id.Value())
		require.True(t, id.IsDebit())
	})
}
