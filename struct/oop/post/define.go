package post

type IPost interface {
	GetTitle() string
	GetContent() string
	GetStatus() string
	SetStatus(status string)
}

type Post struct {
	title   string
	content string
	status  string
}
