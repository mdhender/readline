package main

import (
	"fmt"

	"github.com/mdhender/readline/readline"
)

func main() {
	c := readline.NewReadWriter("> ")
	input, _, err := c.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(c, ">>", input, "<<")
	fmt.Fprintln(c, ">>", string(input), "<<")
	input, _, err = c.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(c, ">>", input, "<<")
	fmt.Fprintln(c, ">>", string(input), "<<")
}
