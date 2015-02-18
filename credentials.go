package gobulksms

import (
	"errors"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetCredential(userName, password string) (*Credential, error) {
	if userName == "" || password == "" {
		return nil, errors.New("Provide username and password")
	}
	var credential Credential
	credential.Username = userName
	credential.Password = password
	return &credential, nil
}
