package post

type VipPost struct {
	IPost
	*Post
	isVip bool
}

func NewVipPost(title, content, status string) IPost {
	return &VipPost{
		Post: &Post{
			title:   title,
			content: content,
			status:  status,
		},
		isVip: false,
	}
}

func (p *VipPost) GetTitle() string {
	return p.title
}

func (p *VipPost) GetContent() string {

	if !p.isVip {
		return "you don't have permissions"
	}

	return p.content
}

func (p *VipPost) GetStatus() string {
	return p.status
}

func (p *VipPost) SetStatus(status string) {
	p.status = status
}
