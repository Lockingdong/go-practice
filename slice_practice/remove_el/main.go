package main

import "fmt"

func main() {

	post := &Post{
		Title:    "test title",
		Content:  "test content",
		Hashtags: []string{"bird", "nature", "animal", "tree", "sun"},
	}

	post.RemoveHashtag("animal")

	post.PrintHashtags() // [bird nature tree sun]
}

type Post struct {
	Title    string
	Content  string
	Hashtags []string
}

func (p *Post) RemoveHashtag(tag string) {
	idx := 0
	for _, v := range p.Hashtags {
		if v != tag {
			p.Hashtags[idx] = v
			idx++
		}
	}

	p.Hashtags = p.Hashtags[:idx]
}

func (p *Post) PrintHashtags() {
	fmt.Println(p.Hashtags)
}
