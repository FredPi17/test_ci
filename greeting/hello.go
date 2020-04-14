package greeting

import "fmt"

func Hello(user string) string {
	if len(user) == 0 {
		return "Hello dude"
	}
	return fmt.Sprintf("Hello %v!", user)
}
