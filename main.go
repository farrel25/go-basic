package main

import (
	"fmt"
	"math"
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

func main() {
	implementationOfInterface()
}
