package main

import "fmt"

func findStudent(id int) string {
	// var temp student
	var isFound bool = false

	if id < 1 {
		return fmt.Sprintln("Input is not valid. id must bigger than 0")
	} else if id > len(students) {
		return fmt.Sprintln("Input is not valid. id must smaller or equal to number of students, which is", len(students))
	}

	for i, v := range students {
		if id-1 == i {
			// temp = v
			isFound = true
			// fmt.Printf("%+v\n", temp)
			fmt.Println("\n\tSTUDENT DATA")
			fmt.Println("id\t\t:", v.id)
			fmt.Println("Name\t\t:", v.name)
			fmt.Println("Address\t\t:", v.address)
			fmt.Println("Occupation\t:", v.occupation)
			fmt.Println("Reason\t\t:", v.reason)
			break
		}
	}

	if !isFound {
		return fmt.Sprintln("Student with id", id, "is not exist in database")
	}

	return ""
}
