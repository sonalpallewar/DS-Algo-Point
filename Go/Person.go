package main

import (
    "fmt"
)

// Define a struct type, which is similar to a class
type Person struct {
    Name string
    Age  int
}

// Method associated with the Person struct (similar to a class method)
func (p *Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// Another struct type
type Employee struct {
    Person
    Position string
    Salary   float64
}

// Method associated with the Employee struct
func (e *Employee) GetDetails() string {
    return fmt.Sprintf("Name: %s\nAge: %d\nPosition: %s\nSalary: $%.2f", e.Name, e.Age, e.Position, e.Salary)
}

func main() {
    // Create an instance of Person
    person := Person{
        Name: "Alice",
        Age:  30,
    }
    fmt.Println(person.Greet())

    // Create an instance of Employee
    employee := Employee{
        Person: Person{
            Name: "Bob",
            Age:  45,
        },
        Position: "Manager",
        Salary:   75000.00,
    }
    fmt.Println(employee.GetDetails())
}
