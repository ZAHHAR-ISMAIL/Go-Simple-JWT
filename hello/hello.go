package main

import (
	"fmt"

	"github.com/ZAHHAR-ISMAIL/Go-Simple-JWT/auth"
)

func main() {
	// Get a greeting message and print it.
	message := auth.Hello("Gladys")
	fmt.Println(message)
}
