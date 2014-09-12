package GoJira

import ()

type Session struct {
	url string
	username  string
	password  string
	useragent string
}

func NewSession(url, username, password, useragent string) *Session {
	return &Session{
		url: url,
		username:  username,
		password:  password,
		useragent: useragent,
	}
}
