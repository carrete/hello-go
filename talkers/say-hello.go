package talkers

import "fmt"

// SayHello says hello.
func SayHello(whom string) string {
	return fmt.Sprintf("Hello, %s!", whom)
}
