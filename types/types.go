package types

type HTTPResponse struct {
	Msg interface{} `json:"Message"`
}

type URL struct {
	URL string `json:"url"`
}
