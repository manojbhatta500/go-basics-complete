package main

import "fmt"

type Grades struct {
	grades_list []int
	name        string
	course      string
}

func main() {

	var changedMarks []int = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	p1 := Grades{grades_list: []int{1, 2, 3, 4, 5}, name: "manoj bhatta", course: "Courses"}
	fmt.Println("before change")
	fmt.Println(p1)
	p1.changeGrades(changedMarks)
	fmt.Println("after change")
	fmt.Println(p1)
}

func (g *Grades) changeGrades(grades []int) {
	g.grades_list = grades
	g.name = "hey ram"
	g.course = "light"
	fmt.Println(g)
}
