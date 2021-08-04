package controllers

type Error struct {
	StatusCode uint
	Message    string
}

func (e Error) Error() string {
	return e.Message
}
