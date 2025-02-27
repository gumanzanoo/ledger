package vo

import (
	"fmt"
	"regexp"
)

var (
	ErrInvalidDocument = fmt.Errorf("invalid document")
)

type Document struct {
	value string
}

func ParseDocument(value string) (Document, error) {
	if value == "" {
		return Document{}, ErrInvalidDocument
	}

	if !regexp.MustCompile(`^\d{11}$|^\d{14}$`).MatchString(value) {
		return Document{}, ErrInvalidDocument
	}
	return Document{value: value}, nil
}

func (d Document) Value() string {
	return d.value
}

func (d Document) IsEmpty() bool {
	return d.value == ""
}

type (
	UserDocument         struct{ Document }
	AccountOwnerDocument struct{ Document }
)

// UserDocument -------------------------------------------------------------------------------------------------------

func ParseUserDocument(value string) (UserDocument, error) {
	doc, err := ParseDocument(value)
	return UserDocument{doc}, err
}

// AccountOwnerDocument -----------------------------------------------------------------------------------------------

func ParseAccountOwnerDocument(value string) (AccountOwnerDocument, error) {
	doc, err := ParseDocument(value)
	return AccountOwnerDocument{doc}, err
}
