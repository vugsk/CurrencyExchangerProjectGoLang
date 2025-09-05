package models

import "time"

type UserModel struct {
	Id         uint
	IdProfile  string
	Login      string
	Name       string
	Email      string
	Password   string
	IdTelegram string
}

type UserSessionModel struct {
	Id         uint
	IdProfile  string
	IdSession  string
	TimeCreate time.Time
	MaxTime    time.Time
}

type RequestChangeCryptModel struct {
	Id          uint
	IdProfile   string
	SumSend     float32
	CryptWallet string
	NumberBank  string
}
