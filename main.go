package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

/* <<< CLOSURE >>> */
// Callback with parameter alias
type isOddNum func(int) bool

func main() {
	/* <<< FUNCTION >>> */
	fmt.Println("\n <<< FUNCTION >>> \n")

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

	/* <<< CLOSURE >>> */
	fmt.Println("\n\n\n <<< CLOSURE >>> \n")

	// Declare closure in variable
	fmt.Println(" >> Declare closure in variable")
	var evenNumbers = func(numbers ...int) []int {
		var result []int
		for _, num := range numbers {
			if num%2 == 0 {
				result = append(result, num)
			}
		}
		return result
	}
	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
	fmt.Println(evenNumbers(numbers...))

	// IIFE (immediately-invoked function expression)
	fmt.Println("\n\n >> IIFE (immediately-invoked function expression)")
	var isPalindrome = func(word string) bool {
		var reversedWord string

		// for i := len(word) - 1; i >= 0; i-- {
		for i := utf8.RuneCountInString(word) - 1; i >= 0; i-- {
			reversedWord += string(word[i])
			fmt.Println("byte:", word[i], ", value:", string(word[i]), ", reversed word:", reversedWord)
		}

		// for _, v := range word {
		// 	fmt.Println("byte:", v, "value:", string(v))
		// }

		return word == reversedWord
	}("katak")
	fmt.Println("is Palindrom?", isPalindrome)

	// Closure as a return value
	fmt.Println("\n\n >> Closure as a return value")
	studentList2 := []string{"Farrel", "Athaillah", "Putra", "Jihan", "Jana"}
	find := findStudent(studentList2)
	fmt.Println(find("PUTRA"))

	// Callback
	// Callback is a closure that is used as a parameter to a function
	fmt.Println("\n\n >> Callback")
	var numbers2 = []int{2, 5, 8, 10, 3, 99, 23}
	var find2 = findOddNumbers(numbers2, func(i int) bool {
		return i%2 != 0
	})
	fmt.Println("Total odd numbers:", find2)

	// Callback with parameter alias
	fmt.Println("\n\n >> Callback with parameter alias")
	var numbers3 = []int{2, 5, 8, 10, 3, 99, 23}
	var find3 = findOddNumbers2(numbers3, func(i int) bool {
		return i%2 != 0
	})
	fmt.Println("Total odd numbers:", find3)

	/* <<< POINTER >>> */
	fmt.Println("\n\n\n <<< POINTER >>> \n")

	// create a pointer
	var number1 *int
	var number2 *int
	_, _ = number1, number2

	// Memory address
	fmt.Println(" >> Memory address")
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println("firstNumber (value)		:", firstNumber)
	fmt.Println("firstNumber (memory address)	:", &firstNumber)
	fmt.Println("secondNumber (value)		:", *secondNumber)
	fmt.Println("secondNumber (memory address)	:", secondNumber)

	// Changing value through pointer
	fmt.Println("\n\n >> Changing value through pointer")
	var firstPerson string = "Farrel"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)		:", firstPerson)
	fmt.Println("firstPerson (memory address)	:", &firstPerson)
	fmt.Println("secondPerson (value)		:", *secondPerson)
	fmt.Println("secondPerson (memory address)	:", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Putra"
	fmt.Println("firstPerson (value)		:", firstPerson)
	fmt.Println("firstPerson (memory address)	:", &firstPerson)
	fmt.Println("secondPerson (value)		:", *secondPerson)
	fmt.Println("secondPerson (memory address)	:", secondPerson)

	// Pointer as a parameter
	fmt.Println("\n\n >> Pointer as a parameter")
	var a int = 10
	fmt.Println("Before (value)		:", a)
	fmt.Println("Before (memory address)	:", &a)

	changeValue(&a)

	fmt.Println("Before (value)		:", a)
	fmt.Println("Before (memory address)	:", &a)
}

/* <<< FUNCTION >>> */
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

/* <<< CLOSURE >>> */

// Closure as a return value
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(s) == strings.ToLower(v) {
				student = v
				position = i + 1
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!!", s)
		}

		return fmt.Sprintf("We found %s at number %d", student, position)
	}
}

// Callback
// Callback is a closure that is used as a parameter to a function
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers += 1
		}
	}

	return totalOddNumbers
}

// Callback with parameter alias
func findOddNumbers2(numbers []int, callback isOddNum) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers += 1
		}
	}

	return totalOddNumbers
}

/* <<< POINTER >>> */

// Pointer as a parameter
func changeValue(number *int) {
	*number = 20
}
