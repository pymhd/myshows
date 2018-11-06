package myshows

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	AuthURL = "https://myshows.me/oauth/token"
	GrantType = "password"
	ContentType = "Content-Type"
	PayloadType = "application/json"
)

type authRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	User         string `json:"username"`
	Password     string `json:"password"`
}

type authResponse struct {
	Token string `json:"access_token"`
}

func GetToken(id, scrt, usr, pwd string) (string, error) {
	ar := authRequest{GrantType, id, scrt, usr, pwd}
	
	pl, _ := json.Marshal(ar)
	plr := bytes.NewReader(pl)
	
	cl := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, AuthURL, plr)
	req.Header.Set(ContentType, PayloadType)

	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	var arp authResponse
	if err := json.NewDecoder(resp.Body).Decode(&arp); err != nil {
                return "", err
        }

	return arp.Token, nil
}
