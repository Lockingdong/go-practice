package post

type GuestPost struct {
	IPost
	*Post
}

func NewGuestPost(title, content, status string) IPost {
	return &GuestPost{
		Post: &Post{
			title:   title,
			content: content,
			status:  status,
		},
	}
}

func (p *GuestPost) GetTitle() string {
	return p.title
}

func (p *GuestPost) GetContent() string {
	return p.content
}

func (p *GuestPost) GetStatus() string {
	return p.status
}

func (p *GuestPost) SetStatus(status string) {
	p.status = status
}
