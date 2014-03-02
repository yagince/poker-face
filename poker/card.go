package poker

type Choice struct {
	Author string `json:"author"`
	Card string `json:"card"`
	Reset bool `json:"reset"`
	Open bool `json:"open"`
}

func (self *Choice) String() string {
	return self.Author + " choices " + self.Card
}

func (self *Choice) ShouldSave() bool {
	return !self.Reset && !self.Open
}
