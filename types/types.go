package types

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
}

type Res struct {
	Quotes []string `json:"data"`
}

type Qt struct {
	Quote []Res `json:"inspirational_quotes"`
}
