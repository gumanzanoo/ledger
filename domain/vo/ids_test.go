package vo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIDs(t *testing.T) {
	t.Parallel()
	const validUUID = "901cbb87-40f3-4c7f-88c3-5419e89195d4"

	t.Run("user id", func(t *testing.T) {
		t.Parallel()

		id, err := ParseUserID(validUUID)
		require.NoError(t, err)
		require.Equal(t, validUUID, id.Value())
		require.False(t, id.IsZero())

		invalidID, err := ParseUserID("")
		assert.ErrorIs(t, err, ErrInvalidUUID)
		assert.True(t, invalidID.IsZero())
	})

	t.Run("account id", func(t *testing.T) {
		t.Parallel()

		id, err := ParseAccountID(validUUID)
		require.NoError(t, err)
		require.Equal(t, validUUID, id.Value())
		require.False(t, id.IsZero())

		invalidID, err := ParseAccountID("")
		assert.ErrorIs(t, err, ErrInvalidUUID)
		assert.True(t, invalidID.IsZero())
	})

	t.Run("transaction id", func(t *testing.T) {
		t.Parallel()

		id, err := ParseTransactionID(validUUID)
		require.NoError(t, err)
		require.Equal(t, validUUID, id.Value())
		require.False(t, id.IsZero())

		invalidID, err := ParseTransactionID("")
		assert.ErrorIs(t, err, ErrInvalidUUID)
		assert.True(t, invalidID.IsZero())
	})

	t.Run("related transaction id", func(t *testing.T) {
		t.Parallel()

		id, err := ParseRelatedTransactionID(validUUID)
		require.NoError(t, err)
		require.Equal(t, validUUID, id.Value())
		require.False(t, id.IsZero())

		invalidID, err := ParseRelatedTransactionID("")
		assert.ErrorIs(t, err, ErrInvalidUUID)
		assert.True(t, invalidID.IsZero())
	})
}
