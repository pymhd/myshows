package myshows

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/pymhd/go-logging"
	_ "io/ioutil"
	"net/http"
)

const (
	Id                 = 1
	JsonRPC            = "2.0"
	ListEpisodesMethod = "lists.Episodes"
	TopShowMethod      = "shows.Top"
	ListShowMethod     = "profile.Shows"
	SearchShowMethod   = "shows.Search"
	ManageShowMethod   = "manage.SetShowStatus"
	GetByIdMethod      = "shows.GetById"
	ApiURL             = "https://api.myshows.me/v2/rpc/"
)

type Request struct {
	Id      int                    `json:"id"`
	JsonRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

var (
	//some predefined params
	paramsListUnwatched = map[string]interface{}{"list": "unwatched"}
)

func GetNextEpisodes(token string) ([]EpisodeDesc, error) {
	var epr EpisodesResponse
	r := Request{Id, JsonRPC, ListEpisodesMethod, paramsListUnwatched}

	resp, err := makeRequst(token, r)
	if err != nil {
		return epr.Result, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&epr)
	if epr.Error.Code != 0 {
		err = fmt.Errorf("%s", epr.Error.Message)
	}

	return epr.Result, err
}

func GetShowList(token string) ([]ShowDesc, error) {
	var sr ShowsResponse
	r := Request{Id, JsonRPC, ListShowMethod, paramsListUnwatched}

	resp, err := makeRequst(token, r)
	if err != nil {
		return sr.Result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&sr)
	if sr.Error.Code != 0 {
		err = fmt.Errorf("%s", sr.Error.Message)
	}

	return sr.Result, err

}

func SearchShow(name string) ([]Show, error) {
	var slr ShowsLookupResponse
	params := map[string]interface{}{"query": name}
	r := Request{Id, JsonRPC, SearchShowMethod, params}

	resp, err := makeRequst("", r)
	if err != nil {
		return slr.Result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&slr)
	if slr.Error.Code != 0 {
		err = fmt.Errorf("%s", slr.Error.Message)
	}

	return slr.Result, err
}

func SetShowAsWatching(token string, id int) error {
	params := map[string]interface{}{"id": id, "status": "watching"}
	r := Request{Id, JsonRPC, ManageShowMethod, params}

	resp, err := makeRequst(token, r)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil

}

func GetTopShows(n int) ([]ShowDesc, error) {
	var tlr TopLookupResponse
	params := map[string]interface{}{"count": n, "mode": "all"}
	r := Request{Id, JsonRPC, TopShowMethod, params}

	resp, err := makeRequst("", r)
	if err != nil {
		return tlr.Result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&tlr)
	if tlr.Error.Code != 0 {
		err = fmt.Errorf("%s", tlr.Error.Message)
	}

	return tlr.Result, err
}

func GetShowById(id int) (Show, error) {
	var gsbir GetShowByIdResponse
	params := map[string]interface{}{"showId": id, "withEpisodes": false}
	r := Request{Id, JsonRPC, GetByIdMethod, params}

	resp, err := makeRequst("", r)
	if err != nil {
		return gsbir.Result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&gsbir)
	if gsbir.Error.Code != 0 {
		err = fmt.Errorf("%s", gsbir.Error.Message)
	}

	return gsbir.Result, err
}

func makeRequst(t string, r Request) (*http.Response, error) {
	pl, _ := json.Marshal(r)
	plr := bytes.NewReader(pl)

	cl := &http.Client{}

	req, _ := http.NewRequest(http.MethodPost, ApiURL, plr)
	req.Header.Set(ContentType, PayloadType)
	req.Header.Set("Authorization", "Bearer "+t)

	resp, err := cl.Do(req)
	if err != nil {
		log.Errorf("Unseccessful api request:  %s\n", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Unseccessful api request. Got %d\n", resp.StatusCode)
		return nil, fmt.Errorf("Network error")
	}

	return resp, nil
}
