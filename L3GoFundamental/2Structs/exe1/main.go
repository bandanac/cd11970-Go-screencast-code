package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
}

type Classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []Student
}

func main() {

	c1 := Classroom{
		id:       1,
		capacity: 1,
		subject:  "sub1",
		studentList: []Student{
			{
				id:   1,
				name: "name1",
			},
			{
				id:   2,
				name: "name2",
			},
		},
	}

	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 2
	c2.subject = "sub2"
	c2.studentList = []Student{
		{
			id:   3,
			name: "name3",
		},
		{
			id:   4,
			name: "name4",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)

	c3 := Classroom{
		id:       1,
		capacity: 200,
		subject:  "Art",
		studentList: []Student{
			{
				id:   20,
				name: "Eric",
			},
			{
				id:   30,
				name: "Sloan",
			},
		},
	}

	c4 := new(Classroom)
	c4.id = 2
	c4.capacity = 100
	c4.subject = "Theater"
	c4.studentList = []Student{
		{
			id:   40,
			name: "Vince",
		},
		{
			id:   50,
			name: "Johnny",
		},
	}

	fmt.Println(c3)
	fmt.Println(c4)

}
