package go_say_hello

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func SayGoodBye(name string) string {
	return fmt.Sprintf("Good Bye, %s", name)
}
