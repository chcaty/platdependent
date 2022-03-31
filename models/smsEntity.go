package main

type SmsEntity struct {
	mobile      string
	signature   string
	templateKey string
	data        map[string]string
	outId       string
}
