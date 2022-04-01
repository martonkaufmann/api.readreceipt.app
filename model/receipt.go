package model

type Receipt struct {
	ID        string
	IsRead    bool
	Recipient string
	Subject   string
	Body      string
	Timestamp int64
}
