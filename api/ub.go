package api

type Part struct {
	Title   string `json:"title"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Authors string `json:"authors"`
}

type Section struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Paper struct {
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	Sections []Section `json:"sections"`
	Author   string    `json:"author"`
}

type UrantiaBook struct {
	Parts  []Part  `json:"parts"`
	Papers []Paper `json:"papers"`
}
