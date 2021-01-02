package main

import (
	"fmt"
	"net/http"
	"sync"
)

type myShows struct {
	t  string
	c  *http.Client
	mu sync.Mutex
}

func (m *myShows) SetToken(token string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.t = token
}

func (m *myShows) Auth(id, secret, user, password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.auth(id, secret, user, password)
}

func (m *myShows) GetUnwatchedEpisodes() ([]EpisodeDesc, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.getUnwatchedEpisodes()
}

func (m *myShows) GetShowsList() ([]ShowDesc, error) {
	m.mu.Lock()
        defer m.mu.Unlock()
        
        return m.getUnwatchedEpisodes
}


func New() *myShows {
	m := new(myShows)
	m.c = httpClient
	return m
}

/*
func main() {
	m := New()
	//m.Auth("myshows_aram808", "CP7Nh2EaGcmVBXnsLv6tyJud", "nemo88", "aram88")
	m.SetToken("0b468b2ddaa554344a2d25b2b82890da3c4be531")
	l, _ := m.GetUnwatchedEpisodes()
	for _, ep := range l {
		fmt.Printf("%+v\n", ep)
	}
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

*/
