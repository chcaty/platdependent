package main

type IOfficialFactory interface {
	GetCode(officialScope OfficialScope) (code string, err error)
	GetUserAccessToken(code string) (accessToken OfficialAccessToken, err error)
	RefreshUserAccessToken(refreshToken string) (accessToken OfficialAccessToken, err error)
}
