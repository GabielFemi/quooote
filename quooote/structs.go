package quooote

type Quote struct {
	Title string
	Body string
	Author interface{}
}

type Quoote struct {
	Id	int
	Punchline string
	Body string
	StartedBy int
	Status int
}