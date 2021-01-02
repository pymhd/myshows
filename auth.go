package myshows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	grantType = "password"
	authUrl   = "https://myshows.me/oauth/token"
)

type authRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	User         string `json:"username"`
	Password     string `json:"password"`
}

type Token struct {
	Expire int    `json:"expires_in"`
	Scope  string `json:"scope"`
	Token  string `json:"access_token"`
}

func (m *myShows) auth(id, secret, user, password string) error {
	ar := authRequest{grantType, id, secret, user, password}

	pl, _ := json.Marshal(ar)
	plr := bytes.NewReader(pl)

	req, _ := http.NewRequest(http.MethodPost, authUrl, plr)
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	var t Token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return err
	}
	fmt.Println(t.Token, t.Scope, t.Expire)
	m.t = t.Token

	return nil
}
