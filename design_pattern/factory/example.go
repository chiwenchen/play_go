package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Times", 65, "Ivan Chen")
	paper1, _ := newPublication("newspaper", "New Yorker", 10, "Ivan Chen")

	fmt.Println(mag1)
	fmt.Println(mag1.getTitle())
	s := mag1.(*magazine)
	s.setFrontPerson("Michael Jordan")
	fmt.Println(s.getFrontPerson())

	fmt.Println(paper1)
	paper1.setPages(65)
	fmt.Println(paper1.getPages())

	// 不好的作法
	s.Price = 990
	fmt.Println(s.Price)
}
