package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var arg = os.Args[1]
	inputId, _ := strconv.Atoi(arg)

	var studentData = findStudent(inputId)
	fmt.Println(studentData)
}
