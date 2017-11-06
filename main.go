package main

import (
	"fmt"

	"github.com/mdhender/readline/readline"
)

func main() {
	c := readline.New("> ")
	input, _ := c.ReadLine()
	fmt.Fprintln(c, ">>", input, "<<")
	input, _ = c.ReadLine()
	fmt.Fprintln(c, ">>", input, "<<")
}
