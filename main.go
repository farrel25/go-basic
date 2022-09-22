package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("Farrel", 22)

	greet2("Farrel", "Semarang")

	var names = []string{"Farrel", "Putra"}
	var printMsg = greet3("Hi", names)
	fmt.Println(printMsg)

	var circumference, area = calculateCircumferenceAndAreasOfCircle(7.5)
	fmt.Println("Circumference:", circumference)
	fmt.Println("Area:", area)

	var circumference2, area2 = calculateCircumferenceAndAreasOfCircle2(7.5)
	fmt.Println("Circumference:", circumference2)
	fmt.Println("Area:", area2)

	studentList := print("Farrel", "Athaillah", "Putra", "Jihan", "Jana")
	fmt.Printf("%v\n", studentList)

	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberList...)
	fmt.Println("Result:", result)

	profile("Farrel", "Rendang", "Fried rice", "Mcd", "KFC")
}

func greet(name string, age int8) {
	fmt.Printf("Hello there! my name is %s, and I'm %d years old", name, age)
}

// function with one parameter but all data types are the same
func greet2(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

// function with return value
func greet3(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	// the Sprintf function will return a value while the Printf function will not
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// function with multiple return value
func calculateCircumferenceAndAreasOfCircle(radius float64) (float64, float64) {
	// calculate areas
	var area float64 = math.Pi * math.Pow(radius, 2)

	// calculate circumference
	var circumference float64 = 2 * math.Pi * radius

	return circumference, area
}

// function with predefined return value
func calculateCircumferenceAndAreasOfCircle2(radius float64) (circumference float64, area float64) {
	// calculate areas
	area = math.Pi * math.Pow(radius, 2)

	// calculate circumference
	circumference = 2 * math.Pi * radius

	return
}

/*
Variadic function
Function with infinite similar parameters
The function can accept an infinite number of arguments
*/
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student-%d", i+1)
		tempMap := map[string]string{
			key: v,
		}
		result = append(result, tempMap)
	}

	return result
}

/*
Variadic function #2
Using the slice data type as a parameter of a variadic function
*/
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

/*
Variadic function #3
Combining regular parameters with variadic parameters in a variadic function.
Variadic parameters need to be put in the final position of the parameter.
*/
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there! My name is", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
