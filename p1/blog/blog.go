package blog

import "fmt"

type BLogInterFace interface {
	GetBlog()
	SetBlog(string, string, string)
}

type Blog struct {
	Title       string
	Description string
	Author      string
}

func New(title string, desc string, author string) BLogInterFace {
	return Blog{title, desc, author}
}

func (b Blog) GetBlog() {
	fmt.Println(b)
}

func (b Blog) SetBlog(title string, desc string, author string) {

}
