package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"
	"time"
)

/*
We can use the shape interface if we have implemented
the methods defined by the shape interface through other
data types. We will implement the methods of the shape
interface via the struct data type.
*/
type shape interface {
	area() float64
	perimeter() float64
}

/*
Since we want to use the shape interface data type and
we want to use it for circle and rectangle structs,
both of these structs need to implement the methods
defined by the shape interface.
*/
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(c.radius, 3)
}

func printAreaAndPerimeter(text string, s shape) {
	fmt.Printf("\n%s area: %.2f\n", text, s.area())
	fmt.Printf("%s perimeter: %.2f\n", text, s.perimeter())
}

func implementationOfInterface() {
	fmt.Println("\n <<< Interface >>> \n")
	/*
		The variable c1 has an interface shape data type,
		and has been assigned a value in the form of
		a struct circle or initialized using struct circle.
		This can happen because the struct circle has
		implemented all the methods defined by the interface shape.
	*/
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	/*
		The results that we see on the terminal are the concrete data
		types or the original data types of the variables c1 and r1.
		The two variables have 2 data types, namely structs that have
		been used as values and interface shapes.

		This is called polymorphism or dynamic type. By implementing
		an interface, a variable can have 2 data types.
	*/
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	// Now let's call the methods we have created
	printAreaAndPerimeter("Circle", c1)
	printAreaAndPerimeter("Rectangle", r1)

	/*
		When the circle struct adds a method other than the method defined
		by the interface shape, then the variable c1 cannot use the new method.
		This happens because the interface shape does not define the volume
		method and the c1 variable we have made into a variable with
		the interface shape data type.
	*/
	// fmt.Printf("Circle volume: %.2f\n", c1.volume()) // ERROR : c1.volume undefined (type shape has no field or method volume)
	/*
		Type assertion
		To overcome this error, a type assertion technique is needed, where
		this technique serves to return an interface data type to its original
		data type. In our case this time, we want to return the variable c1
		to its original data type, which is the struct circle data type.
	*/
	fmt.Printf("\nCircle Volume: %.2f\n", c1.(circle).volume())

	// checks whether a type assertion that we use is successful or not
	value, ok := c1.(circle)
	fmt.Println("\nValue:", value, "|", "Ok:", ok)

	if ok {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %.2f\n", value.volume())
	}
}

// =====================================================================

func implementationOfEmptyInterface() {
	/*
		Empty interface is a data type that can accept any data type
		in the Go language. Therefore, a variable with data type empty
		interface can be assigned values with different data types.
	*/
	var randomValues interface{}
	_ = randomValues
	randomValues = "Semarang"
	randomValues = 22
	randomValues = true
	randomValues = []string{"Farrel", "Putra"}

	// Type assertion
	var v interface{}
	v = 20
	// v *= 9 // ERROR : invalid operation: v *= 9 (mismatched types interface{} and int)
	/*
		we can only multiply the concrete or original int data type,
		while the variable v has an empty interface data type
		which is assigned a value of int data type.

		To overcome the error, then we can do a type assertion
	*/
	if value, ok := v.(int); ok {
		v = value * 9
	}

	// Map & Slice with Empty Interface
	/*
		When we declare a map and its values are assigned
		an empty interface data type, then the map can have
		values with different data types. Similar to slices
		or arrays, when they are declared and their contents
		are given the data type empty interface, the slices
		or arrays can have values with different data types.
	*/
	rs := []interface{}{1, "Farrel", true, 2.4, "Putra", false}
	rm := map[string]interface{}{
		"name":   "Farrel",
		"ntatus": true,
		"age":    22,
	}
	_, _ = rs, rm
}

// =====================================================================

func implementationOfReflect() {
	fmt.Println("\n\n\n <<< Reflect >>> \n")
	/*
		Reflect is used to inspect variables, retrieve information from them or even manipulate them.
	*/

	// Identifying Data Type & Value
	fmt.Println("\n >>> Identifying Data Type & Value \n")
	var number = 23
	var reflectValue = reflect.ValueOf(number) // returns an object of type reflect.Value, which contains information about the corresponding variable

	fmt.Println("Reflect Value:", reflectValue)
	fmt.Println("Variable type:", reflectValue.Type()) // returns the data type of the variable in the form of a string

	// To display the reflect variable value, we must first make sure the data type
	// Check the type of data type using the Kind() method
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("\nreflectValue.Kind():", reflectValue.Kind())
		fmt.Println("reflect.Int:", reflect.Int)
		fmt.Println("Variable value:", reflectValue.Int())
	}

	// Accessing Value Using Interface Method
	fmt.Println("\n\n >>> Accessing Value Using Interface Method \n")

	fmt.Println("Variable value:", reflectValue.Interface()) // If the value is only required to be displayed to the output, you can use .Interface().

	/*
		The Interface() method returns an empty interface value or interface{}.
		The original value itself can be accessed by casting the empty interface.
	*/
	var val = reflectValue.Interface().(int)
	fmt.Println("Variable value:", val)
}

// =====================================================================

func goroutine() {
	fmt.Println("Hello")
}

// Asynchronous process #1
func firstProcess(number int) {
	fmt.Println("firstProcess function started")
	for i := 1; i <= number; i++ {
		fmt.Println("i =", i)
	}
	fmt.Println("firstProcess function ended")
}

func secondProcess(number int) {
	fmt.Println("secondProcess function started")
	for j := 1; j <= number; j++ {
		fmt.Println("j =", j)
	}
	fmt.Println("secondProcess function ended")
}

func implementationOfGoroutines() {
	fmt.Println("\n\n\n <<< Concurrency & Goroutines >>> \n")
	/*
		Goroutines are asynchronous, so the process does
		not wait for each other with other Goroutines
	*/

	// go goroutine()

	// Asynchronous process #1
	fmt.Println("\n >>> Asynchronous process #1 \n")
	fmt.Println("Main execution started")

	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 3)

	fmt.Println("Main execution ended")
	/*
		If we pay attention, the firstProcess function does not
		display the result. This happens because each Goroutine
		works asynchronously (it means it won't wait until the
		goroutine is finished) and one Goroutine will not wait
		on another Goroutine.

		there are 2 total goroutines running even though we only
		run one function that is used as a goroutine. This happens
		because of the fact that the main function is also a Goroutine,
		so the main function will not wait for another Goroutine
		to finish executing. This is the cause of the firstProcess
		function not displaying the results even though the function
		has actually been executed.
	*/
}

// =====================================================================

// Waitgroup is a struct of package sync, which is used to synchronize goroutines.

/*
The waitgroup in the printFruit function is a pointer,
this needs to be done so that the waitgroup in the main
and printFruit functions contains the same memory.
*/
func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s\n", index, fruit)
	wg.Done() // to notify the waitgroup about the goroutine that has completed its processing
}

func implementationOfWaitgroup() {
	fruits := []string{"apple", "mango", "durian", "avocado"}

	var wg sync.WaitGroup

	/*
		The Add method is used to add the counter of the waitgroup.
		The waitgroup counter is useful for informing the waitgroup
		about the number of go routines to wait. Because we looped
		4 times, it means that there are 4 go routines to wait for
		before the main function stops executing
	*/
	for i, fruit := range fruits {
		wg.Add(1)
		go printFruit(i, fruit, &wg)
	}

	/*
		The Wait method is useful for waiting for all the go routines
		to complete their process, so the Wait method will hold the
		main function until the entire go routine process is complete.
	*/
	wg.Wait()
}

func main() {
	// implementationOfInterface()
	// implementationOfEmptyInterface()
	// implementationOfReflect()
	// implementationOfGoroutines()
	implementationOfWaitgroup()
}
