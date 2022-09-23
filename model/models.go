package model

type Work struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Url           string `json:"url"`
	Backgroundimg string `json:"backgroundimg"`
	Usageimg      string `json:"usageimg"`
	Sentence      string `json:"sentence"`
}

type News struct {
	Id       int    `json:id`
	Title    string `json:"title"`
	Topimg   string `json:"topimg"`
	Sentence string `json:"sentence"`
}

type Tackle struct {
	Id string `json:id`
}
