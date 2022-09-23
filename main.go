package main

// import "fmt"
// import "strings"
import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	/*
	   comment code
	   produce Hello World output in command line/terminal
	*/
	fmt.Println("Hello World")
	fmt.Println("Hello", " World!", "How", "are", "you")

	// ftm.Println("This line wil not be execute")

	fmt.Print("Farrel Putra\n")
	fmt.Print("Farrel", " Putra\n")
	fmt.Print("Farrel", " ", "Putra\n")

	/*
	   VARIABLES
	*/
	fmt.Println("\n\n")
	var name string = "Farrel"
	var age int = 20
	var name2 string

	age = 22
	name2 = "Putra"

	fmt.Println("My name is", name, ", I'm", age, "years old")
	fmt.Println("My name is", name2, ", I'm", age, "years old")

	// type inference / variable without data type
	var name3 = "Atha"
	var age3 = 22

	fmt.Printf("%T, %T\n", name3, age3)

	// Short Declaration
	name4 := "Athaillah"
	age4 := 22
	fmt.Println("My name is", name4, ", I'm", age4, "years old")

	// Multiple Variable Declaration
	var student1, student2, student3 string = "one", "two", "three"
	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	// Multiple Variable Declaration with Type Inference
	var name5, age5, address5 = "Farrel", 22, "Tembalang"
	first2, second2, third2 := "1", 2, "3"

	fmt.Println(name5, age5, address5)
	fmt.Println(first2, second2, third2)

	// Underscore Variable
	var firstVariable string
	var name6, age6, addr6 = "Farrel", 22, "Semarang"

	_, _, _, _ = firstVariable, name6, age6, addr6

	// fmt.Prinf Usage
	first3, second3 := "1", 2
	fmt.Printf("Data type of first3 variable: %T \n", first3)
	fmt.Printf("Data type of second3 variable: %T \n", second3)

	name7, age7, addr7 := "Farrel", 22, "Semarang"
	fmt.Printf("Hello, my name is %s. I'm %d years old. I live in %s\n", name7, age7, addr7)

	/*
	   DATA TYPES
	*/
	fmt.Println("\n\n")

	// Number (integer)
	var first4 uint8 = 89
	var second4 int8 = -12
	fmt.Printf("Data type of first4 = %T\n", first4)
	fmt.Printf("Data type of second4 = %T\n", second4)

	// Number (float)
	var decimalNum float32 = 3.64

	fmt.Printf("Data type of decimalNum = %T\n", decimalNum)
	fmt.Printf("Decimal Number = %f\n", decimalNum)
	fmt.Printf("Decimal Number = %.3f\n", decimalNum)
	fmt.Println("Decimal Number =", decimalNum)

	// Bool
	var condition bool = true
	fmt.Printf("Data type of condition variable = %T\n", condition)
	fmt.Printf("Is it permitted ? %t\n", condition)
	fmt.Println("Is it permitted ?", condition)

	// String
	var msg = `Welcome to "Naratik".
  Nice to meet you.
  Let's learn "Scalable Web Service with Go"`
	fmt.Println(msg)

	/*
	   CONSTANTS & OPERATORS
	*/
	fmt.Println("\n\n")

	// Constant
	const full_name string = "Farrel Athaillah Putra"
	fmt.Printf("Hello %s\n", full_name)

	// Operators (Arithmetic Operators)
	var value = 2 + 2*3
	fmt.Println(value)

	var value2 = (2 + 2) * 3
	fmt.Println(value2)

	// Operators (Relational Operators)
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("First Condition:", firstCondition)
	fmt.Println("Second Condition:", secondCondition)
	fmt.Println("Third Condition:", thirdCondition)
	fmt.Println("Fourth Condition:", fourthCondition)

	// Operators (Logical Operators)
	var right = true
	var wrong bool = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)

	/*
	   CONDITIONS
	*/
	fmt.Println("\n\n")

	// Temporary Variable
	var currentYear = 2022
	if age := currentYear - 1998; age < 17 {
		fmt.Println("You are not allowed to make a sim card")
	} else {
		fmt.Println("You are allowed to make a sim card")
	}

	// Switch
	var score = 6
	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}

	// Switch with Relational Operations
	// var score2 = 5

	// switch score2 {
	// case score2 == 8:
	//   fmt.Println("Perfect")
	// case score2 < 8 && score > 3:
	//   fmt.Println("Not bad")
	// default:
	//   {
	//     fmt.Println("Study harder")
	//     fmt.Println("You need to learn more")
	//   }
	// }

	// Switch (fallthrough keyword)
	var score3 = 8
	switch score3 {
	case 8:
		fmt.Println("Perfect")
		fallthrough
	case 7, 6, 5:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("You need to learn more")
		}
	}

	// Nested Conditions
	var score4 = 0

	if score4 > 7 {
		switch score4 {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if score4 == 5 {
			fmt.Println("Not bad")
		} else if score4 == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if score4 == 0 {
				fmt.Println("Try harder!")
			}
		}
	}

	/*
	   LOOPINGS
	*/
	fmt.Println("\n\n")

	// First Way of Looping
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i+1)
	}

	// Second Way of Looping (just like using while)
	var i = 0
	for i < 3 {
		fmt.Println("Angka", i+1)
		i++
	}

	// Third Way of Looping
	var j = 0
	for j < 3 {
		fmt.Println("Angka", j+1)
		j++
		if j == 3 {
			break
		}
	}

	// break and continue keywords
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Nested Looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}

	// Label
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Loop -", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}

	/*
	   ARRAY
	*/
	fmt.Println("\n\n")

	// Array
	var number [4]int
	number = [4]int{1, 2, 3, 4}

	var names = [3]string{"Farrel", "Athaillah", "Putra"}

	fmt.Printf("%#v \n", number)
	fmt.Printf("%v \n", number)
	fmt.Printf("%#v \n", names)
	fmt.Printf("%v \n", names)

	// Modify Element Through Index
	var fruits = [3]string{"Apple", "Banana", "Mango"}
	fruits[0] = "Strawberry"
	fruits[1] = "Avocado"

	fmt.Printf("%#v \n", fruits)

	// Loop Through Element (range loop)
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// Loop Through Element (for loop)
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}

	// Multidimensional Array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, val := range arr {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

	/*
	   SLICE
	*/
	fmt.Println("\n\n")

	// Slice
	var fruits2 = []string{"apple", "banana", "mango"}
	fmt.Println(fruits2)
	fmt.Printf("%#v\n", fruits2)

	// make function
	var fruits3 = make([]string, 3)
	fmt.Printf("%#v\n", fruits3)

	// append function
	// fruits3[0] = "orange"
	// fruits3[1] = "pineapple"
	// fruits3[2] = "avocado"
	fruits3 = append(fruits3, "orange", "pineapple", "avocado")

	fmt.Printf("%#v\n", fruits3)

	// append function with ellipsis (...)
	var fruits4 = []string{"starfruit", "watermelon"}
	fruits3 = append(fruits3, fruits4...)

	fmt.Printf("%#v\n", fruits3)

	// copy function
	var fruits5 = []string{"blackberry", "blackcurrant", "blueberry"}
	var fruits6 = []string{"cherry", "coconut", "durian"}

	nn := copy(fruits5, fruits6)

	fmt.Println("Fruits5 =>", fruits5)
	fmt.Println("Fruits6 =>", fruits6)
	fmt.Println("Copied elements =>", nn)

	// Slicing
	var fruits7 = []string{"grape", "guava", "plum", "kiwi", "lemon"}
	fmt.Printf("fruits7 => %#v\n", fruits7)

	var fruits71 = fruits7[1:4]
	fmt.Printf("fruits7[1:4] => %#v\n", fruits71)

	var fruits72 = fruits7[0:]
	fmt.Printf("fruits7[0:] => %#v\n", fruits72)

	var fruits73 = fruits7[:3]
	fmt.Printf("fruits7[:3] => %#v\n", fruits73)

	var fruits74 = fruits7[:]
	fmt.Printf("fruits7[:] => %#v\n", fruits74)

	// Combining slicing and append
	var fruits8 = []string{"lime", "lychee", "mango", "melon", "orange"}
	fruits8 = append(fruits8[:3], "papaya")
	fmt.Printf("fruits8 => %#v\n", fruits8)

	// Backing array
	var fruits9 = []string{"pear", "pineapple", "raspberry", "redcurrant", "salak"}

	var fruits91 = fruits9[2:4]

	fruits91[0] = "rambutan"

	fmt.Println("Fruits9 =>", fruits9)
	fmt.Println("Fruits91 =>", fruits91)

	// cap function
	var fruits10 = []string{"starfruit", "strawberry", "yuzu", "apple"}
	fmt.Println("Fruits10:", fruits10)
	fmt.Println("Fruits10 len:", len(fruits10))
	fmt.Println("Fruits10 cap:", cap(fruits10))

	fmt.Println(strings.Repeat("#", 25))

	var fruits102 = fruits10[0:3]
	fmt.Println("Fruits102:", fruits102)
	fmt.Println("Fruits102 len:", len(fruits102))
	fmt.Println("Fruits102 cap:", cap(fruits102))

	fmt.Println(strings.Repeat("#", 25))

	var fruits103 = fruits10[1:]
	fmt.Println("Fruits103:", fruits103)
	fmt.Println("Fruits103 len:", len(fruits103))
	fmt.Println("Fruits103 cap:", cap(fruits103))

	// Creating a new backing array
	cars := []string{"Ford", "Honda", "Audi", "RangeRover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println("Cars:", cars)
	fmt.Println("New Cars:", newCars)

	/*
	   MAP
	*/
	fmt.Println("\n\n")

	// map
	var person map[string]string // declaration
	person = map[string]string{} // initialization

	person["name"] = "Farrel"
	person["age"] = "22"
	person["address"] = "Semarang"

	fmt.Println("Name:", person["name"])
	fmt.Println("Age:", person["age"])
	fmt.Println("Address:", person["address"])

	var person2 map[string]string // declaration
	person2 = map[string]string{
		"name":    "Putra",
		"age":     "23",
		"address": "Semarang",
	} // initialization

	fmt.Println("Name:", person2["name"])
	fmt.Println("Age:", person2["age"])
	fmt.Println("Address:", person2["address"])

	// Looping with map
	for k, v := range person {
		fmt.Println(k, ":", v)
	}

	// Deleting value
	fmt.Println("Before deleting:", person)
	delete(person, "age")
	fmt.Println("After deleting:", person)

	// Detecting a value
	value3, exist := person2["age"]

	if exist {
		fmt.Println(value3)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person2, "age")

	value3, exist = person2["age"]

	if exist {
		fmt.Println(value3)
	} else {
		fmt.Println("Value has been deleted")
	}

	// Combining slice with map
	var people = []map[string]string{
		{"name": "Farrel", "age": "21"},
		{"name": "Atha", "age": "22"},
		{"name": "Putra", "age": "23"},
	}

	for idx, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", idx, person["name"], person["age"])
	}

	/*
	   ALIASE
	*/
	fmt.Println("\n\n")

	// variable declaration with uint8 data type
	var a uint8 = 10
	var b byte // byte is an alias of the uint8 data type

	b = a // no error occurs, because byte has the same data type as uint8
	_ = b

	// Declare the alias data type with a name that we create ourselves
	// type alias_name = data_type_name
	type second5 = uint

	var hour second5 = 3600
	fmt.Printf("hour type: %T\n", hour)

	/*
	   Strings In Depth
	*/
	fmt.Println("\n\n")

	// Looping Over String (byte-by-byte)
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}
	for i, v := range city {
		fmt.Println("index:", i, ", value:", string(v), "byte:", v)
	}

	var city2 []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(city2)
	fmt.Println(string(city2))

	// Looping Over String (rune-by-rune)
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("\nstr1 %s byte length => %d\n", str1, len(str1))
	for i := 0; i < len(str1); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, str1[i])
	}

	fmt.Printf("\nstr2 %s byte length => %d\n", str2, len(str2))
	for i := 0; i < len(str2); i++ {
		fmt.Printf("index: %d, byte: %d, string: %s\n", i, str2[i], string(str2[i]))
	}

	// When we want to find the number of characters, and not the number of bytes,
	// then we need to convert the string into runes first.
	fmt.Printf("\nstr1 %s character length => %d\n", str1, utf8.RuneCountInString(str1))
	fmt.Printf("str2 %s character length => %d\n", str2, utf8.RuneCountInString(str2))

	for i, s := range str2 {
		fmt.Printf("index: %d, string: %s, byte: %d\n", i, string(s), s)
	}
}
