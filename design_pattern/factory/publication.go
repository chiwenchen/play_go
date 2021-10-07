package main

// 定義工廠回傳的實例需要有的function
type iPublication interface {
	setTitle(title string)
	setPages(pages int)
	// setPublisher(publisher string)
	getTitle() string
	getPages() int
	// getPublisher() string
}

type publication struct {
	title     string
	pages     int
	publisher string
	Price     int
	// 用 public value可以讓 type assert之後的實例可以直接取用，但是是一個不好的方法，
	// 比較好的方式是改成小寫，並且定義 func來做存取，就像 title一樣
}

// setTitle sets the title, share amoung all types
func (p *publication) setTitle(title string) {
	p.title = title
}

// getTitle gets the title, share amoung all types
func (p *publication) getTitle() string {
	return p.title
}

// 共用的func可以定義在被 embad的 type上
func (p *publication) getPages() int {
	return p.pages
}
