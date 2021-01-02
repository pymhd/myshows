package myshows

type Show struct {
	ID    int    `json:"id"`
	Title string `json:"titleOriginal"`
	Image string `json:"image"`
}

type Episode struct {
	ID         int    `json:"id"`
	ShowID     int    `json:"showId"`
	SeasonNum  int    `json:"seasonNumber"`
	EpisodeNum int    `json:"episodeNumber"`
	Title      string `json:"title"`
	Image      string `json:"image"`
}

type EpisodeDesc struct {
	Show    Show    `json:"show"`
	Episode Episode `json:"episode"`
}

type ShowDesc struct {
	Show Show `json:"show"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type EpisodeResponse struct {
	Result []EpisodeDesc `json:"result"`
	Error  Error         `json:"error"`
}

type ShowResponse struct {
	Result []ShowDesc `json:"result"`
	Error  Error      `json:"error"`
}
