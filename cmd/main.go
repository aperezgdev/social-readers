package main

import (
	"github.com/aperezgdev/social-readers-api/cmd/boostrap"
)

func main() {
	if err := boostrap.Run(); err != nil {
		panic(err)
	}
}
