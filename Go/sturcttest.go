type Greeter interface {
    Greet() string
}

type Person struct {
    Name string
    Age  int
}

func (p *Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}
