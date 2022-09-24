package main

import (
	"fmt"

	"github.com/go-practice/struct/oop/post"
)

func main() {

	var guestPost post.IPost
	guestPost = post.NewGuestPost(
		"test_title",
		"test_content",
		"published",
	)
	status := guestPost.GetStatus()
	content := guestPost.GetContent()
	fmt.Println(status)  // published
	fmt.Println(content) // test_content

	var vipPost post.IPost
	vipPost = post.NewVipPost(
		"test_title",
		"test_content",
		"published",
	)
	status2 := vipPost.GetStatus()
	content2 := vipPost.GetContent()
	fmt.Println(status2)  // published
	fmt.Println(content2) // you don't have permissions
}
