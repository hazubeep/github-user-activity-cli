package model

type GithubEvent struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action  string `json:"action"`
	RefType string `json:"ref_type"`
}
