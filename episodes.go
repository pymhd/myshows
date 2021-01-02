package myshows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

type Params map[string]interface{}

func (m *myShows) getUnwatchedEpisodes() ([]EpisodeDesc, error) {
	method := "lists.Episodes"
	params := map[string]interface{}{"list": "unwatched"}

	el := new(EpisodeResponse)
	err := m.makeQuery(http.MethodPost, method, params, &el)

	return el.Result, err
}

func (m *myShows) makeQuery(httpMethod, method string, params Params, result interface{}) error {
	// enc req body
	data := Request{1, "2.0", method, params}
	payloadBytes, _ := json.Marshal(data)
	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest(httpMethod, "https://api.myshows.me/v2/rpc/", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	authHeader := fmt.Sprintf("Bearer %s", m.t)
	req.Header.Set("Authorization", authHeader)

	resp, err := m.c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}
