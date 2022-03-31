package main

type IQiNiuFactory interface {
	GetToken() (token string, err error)
}
