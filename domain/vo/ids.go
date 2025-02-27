package vo

import (
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrInvalidUUID = fmt.Errorf("invalid uuid")
)

type uuidID struct {
	value uuid.UUID
}

func parseUUID(s string) (uuidID, error) {
	uuidParsed, err := uuid.Parse(s)
	if err != nil {
		return uuidID{}, ErrInvalidUUID
	}

	return uuidID{uuidParsed}, nil
}

func (id uuidID) IsZero() bool {
	return id.value == uuid.Nil
}

func (id uuidID) Value() string {
	if id.IsZero() {
		return ""
	}

	return id.value.String()
}

func (id uuidID) String() string {
	return id.value.String()
}

func (id uuidID) UUID() uuid.UUID {
	return id.value
}

func (id uuidID) IDMarshalJSON() ([]byte, error) {
	return []byte(`"` + id.String() + `"`), nil
}

type (
	UserID               struct{ uuidID }
	AccountID            struct{ uuidID }
	TransactionID        struct{ uuidID }
	RelatedTransactionID struct{ uuidID }
)

// UserID -------------------------------------------------------------------------------------------------------------

func NewUserID() UserID {
	return UserID{uuidID{uuid.New()}}
}

func ParseUserID(s string) (UserID, error) {
	id, err := parseUUID(s)
	if err != nil {
		return UserID{}, err
	}
	return UserID{id}, err
}

// AccountID ----------------------------------------------------------------------------------------------------------

func NewAccountID() AccountID {
	return AccountID{uuidID{uuid.New()}}
}

func ParseAccountID(s string) (AccountID, error) {
	id, err := parseUUID(s)
	if err != nil {
		return AccountID{}, err
	}
	return AccountID{id}, err
}

// TransactionID ------------------------------------------------------------------------------------------------------

func NewTransactionID() TransactionID {
	return TransactionID{uuidID{uuid.New()}}
}

func ParseTransactionID(s string) (TransactionID, error) {
	id, err := parseUUID(s)
	if err != nil {
		return TransactionID{}, err
	}
	return TransactionID{id}, err
}

// TransactionID ------------------------------------------------------------------------------------------------------

func NewRelatedTransactionID() RelatedTransactionID {
	return RelatedTransactionID{uuidID{uuid.New()}}
}

func ParseRelatedTransactionID(s string) (RelatedTransactionID, error) {
	id, err := parseUUID(s)
	if err != nil {
		return RelatedTransactionID{}, err
	}
	return RelatedTransactionID{id}, err
}
