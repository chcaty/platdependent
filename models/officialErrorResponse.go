package main

type OfficialErrorResponse struct {
	errCode int    `json:"errcode"`
	errMsg  string `json:"errmsg"`
}
