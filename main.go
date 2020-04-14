package main

import (
	"fmt"
	"github.com/terraform-providers/test_ci/greeting"
)

func main() {
	greetMessage := greeting.Hello("Michel")

	fmt.Println(greetMessage)
}
