package main

import "fmt"

type Student struct {
	Name  string
	age   int
	marks int
}

func main() {
	students := []Student{
		{
			Name:  "Aditya",
			age:   18,
			marks: 90,
		},
		{
			Name:  "tectoro",
			age:   12,
			marks: 87,
		},
	}
	var b int = 1
	for b == 1 {
		fmt.Println("Student Manager: ")
		var x int
		fmt.Println("1.Add\n2.Delete\n3.Update\n4.Display\n")
		fmt.Scan(&x)

		switch x {
		case 1:
			fmt.Println("Enter the Students Details")
			var S Student
			fmt.Println("Enter the Student Name,Age and MArks: ")
			fmt.Scan(&S.Name)
			fmt.Scan(&S.age)
			fmt.Scan(&S.marks)
			students = append(students, S)
		case 2:
			fmt.Println("Enter the student Name to delete")
			var name string
			fmt.Scan(&name)
			for i, v := range students {
				if v.Name == name {
					students = append(students[:i], students[i+1:]...)
				}
			}
		case 4:
			for i, v := range students {
				fmt.Println(i, "Student", v.age, v.marks, v.age)
			}

		}
		fmt.Println("If u want to continue press(0)")
		var c int
		fmt.Scan(&c)
		if c == 0 {
			b = 0
		}

	}
}
