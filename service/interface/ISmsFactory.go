package main

type ISmsFactory interface {
	Send(smsEntity SmsEntity) (err error)
}
