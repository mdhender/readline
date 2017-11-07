package main

import (
	"fmt"
	"strings"

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
	text, err := c.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text = strings.TrimRight(text, "\r\n")
	fmt.Fprintln(c, ">>", text, "<<")
	text, err = c.ReadToEOL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(c, ">>", text, "<<")
}
