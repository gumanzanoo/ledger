package entities

import (
	"transactions/domain/vo"
)

type User struct {
	id       vo.UserID
	name     string
	document vo.UserDocument
}

func NewUser(name string, document vo.UserDocument) User {
	return User{
		id:       vo.NewUserID(),
		name:     name,
		document: document,
	}
}

func (u User) ID() vo.UserID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Document() vo.UserDocument {
	return u.document
}
