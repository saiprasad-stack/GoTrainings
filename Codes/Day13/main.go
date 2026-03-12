// // package main

// // import "fmt"

// // type Greeter interface {
// // 	Greet() string
// // }

// // type Person struct {
// // 	Name string
// // }

// // func (p Person) Greet() string {
// // 	return "Hello " + p.Name
// // }
// // func main() {
// // 	// Interfaces -- MYSQL , MongoDB , CockrochDB

// // 	var g Greeter
// // 	// p := Person{Name: "Golang"}

// // 	g = Person{Name: "Golang"}
// // 	fmt.Println(g.Greet())

// // }

// package main

// import (
// 	"fmt"
// )

// // Notifier is a basic interface for sending notifications
// // Pattern inspired by Docker's logging drivers

// // Creating an interface
// type Notifier interface {
// 	Notify(message string) error
// 	GetType() string
// }

// // EmailNotifier implements the Notifier interface
// type EmailNotifier struct {
// 	EmailAddress string
// }

// func (e *EmailNotifier) Notify(message string) error {
// 	fmt.Printf("[Email] Sending to %s: %s\n", e.EmailAddress, message)
// 	return nil
// }

// func (e *EmailNotifier) GetType() string {
// 	return "email"
// }

// func main() {
// 	// Create different notifiers
// 	email := &EmailNotifier{EmailAddress: "admin@example.com"}
// 	// slack := &SlackNotifier{WebhookURL: "https://hooks.slack.com/xxx", Channel: "alerts"}
// 	// sms := &SMSNotifier{PhoneNumber: "+1234567890"}

// 	// fmt.Println("=== Basic Interface Example ===")

// 	// // Use different notifiers with the same function
// 	// AlertSystem(email, "Server is running low on memory")
// 	// AlertSystem(slack, "Deployment completed successfully")
// 	// AlertSystem(sms, "Critical: Database is down")

// 	// Store multiple notifiers in a slice
// 	notifiers := []Notifier{email}

//		fmt.Println("\nBulk notifications:")
//		for _, n := range notifiers {
//			n.Notify("System health check passed")
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // Reader defines read behavior
// type Reader interface {
// 	Read() (string, error)
// }

// // Writer defines write behavior
// type Writer interface {
// 	Write(data string) error
// }

// // ReadWriter composes Reader and Writer
// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// // Closer defines close behavior
// type Closer interface {
// 	Close() error
// }

// // ReadWriteCloser composes multiple interfaces
// type ReadWriteCloser interface {
// 	ReadWriter
// 	Closer
// }

// // File represents a simple file (like in os package)
// type File struct {
// 	Name    string
// 	content []string
// 	closed  bool
// }

// func NewFile(name string) *File {
// 	return &File{
// 		Name:    name,
// 		content: make([]string, 0),
// 	}
// }

// func (f *File) Read() (string, error) {
// 	if f.closed {
// 		return "", fmt.Errorf("file is closed")
// 	}
// 	return strings.Join(f.content, "\n"), nil
// }

// func (f *File) Write(data string) error {
// 	if f.closed {
// 		return fmt.Errorf("file is closed")
// 	}
// 	f.content = append(f.content, data)
// 	return nil
// }

// func (f *File) Close() error {
// 	f.closed = true
// 	fmt.Printf("Closing file: %s\n", f.Name)
// 	return nil
// }

// // ProcessData accepts any Reader
// func ProcessData(r Reader) {
// 	data, _ := r.Read()
// 	fmt.Printf("Processing data: %s\n", data)
// }

// // SaveData accepts any Writer
// func SaveData(w Writer, data string) {
// 	w.Write(data)
// 	fmt.Println("Data saved")
// }

// func main() {
// 	fmt.Println("=== Interface Composition ===")

// 	// File implements ReadWriteCloser
// 	file := NewFile("config.txt")

// 	// Use as Writer
// 	SaveData(file, "host=localhost")
// 	SaveData(file, "port=8080")

// 	// Use as Reader
// 	ProcessData(file)

// 	// Use as Closer
// 	file.Close()

// 	// Demonstrate interface satisfaction
// 	var rw ReadWriter = file
// 	rw.Write("new data")
// 	rw.Read()

// 	fmt.Printf("\nFile implements multiple interfaces!\n")
// }

package main

import "fmt"

type Engine interface {
	Start()
}

type PetrolEngine struct{}

func (PetrolEngine) Start() {
	fmt.Println("Petrol Engine Started")
}

type ElectricEngine struct{}

func (ElectricEngine) Start() {
	fmt.Println("Electric Engine Started")

}

type Car struct {
	engine Engine
}

func (c Car) Start() {
	c.engine.Start()
}

func main() {
	petrolCar := Car{engine: PetrolEngine{}}
	petrolCar.Start()

	ECar := Car{engine: ElectricEngine{}}
	ECar.Start()
}

// interface A -- Sing() , Dance() , Talent()

// interface B -- Speak() , Laugh()

// Compostiton -- interface C - A , B , Register()

// There is no implement keyword
