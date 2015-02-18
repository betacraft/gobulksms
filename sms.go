package main

import (
	"bulksms/constants"
	"bytes"
	"errors"
	"net/http"
	"net/url"
)

// Right now supports only text messages
// later can be extended to support various formats

type SMS struct {
	Credentials *Credential
	Msg         *Message
}

func (sms *SMS) Send() error {
	if len(sms.Msg.Text) <= 0 {
		return errors.New("Invalid message text")
	}
	if len(sms.Msg.To) <= 0 {
		return errors.New("Please provide numbers to send")
	}
	params := make(url.Values)
	params.Set("username", sms.Credentials.Username)
	params.Set("password", sms.Credentials.Password)
	params.Set("message", sms.Msg.Text)
	params.Set("msisdn", sms.Msg.To)

	client := &http.Client{}
	urlParams := bytes.NewBufferString(params.Encode())
	var r *http.Request
	r, _ = http.NewRequest("POST", constants.SEND_SMS_URL, urlParams)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err := client.Do(r)
	if err != nil {
		return err
	}
	return nil
}
