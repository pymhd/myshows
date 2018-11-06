package myshows

type Show struct {
	Id            int    `json:"id"`
	TitleOriginal string `json:"titleOriginal"`
	Title         string `json:"title"`
	Image         string `json:"image"`
	Rating        float64 `json:"rating"`
}

type Episode struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	SeasonNum  int    `json:"seasonNumber"`
	EpisodeNum int    `json:"episodeNumber"`
}

type EpisodeDesc struct {
	Show    `json:"show"`
	Episode `json:"episode"`
}

type ShowDesc struct {
	Show `json:"show"`
}

type Error struct {
	Code 	int `json:"code"`
	Message string `json:"message"`
}

type EpisodesResponse struct {
	Result []EpisodeDesc `json:"result"`
	Error    Error       `json:"error"`
}

type ShowsResponse struct {
	Result []ShowDesc `json:"result"`
	Error    Error       `json:"error"`
}

type ShowsLookupResponse struct {
        Result []Show      `json:"result"`
        Error    Error       `json:"error"`
}

type TopLookupResponse struct {
        Result []ShowDesc  `json:"result"`
        Error    Error       `json:"error"`
}
