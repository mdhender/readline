package readline

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Readline should satisfy a ReadWriter interface
type Readline struct {
	rdr        *bufio.Reader
	wrt        io.Writer
	prompt     string
	isTerminal bool
}

// New creates a new Readline by opening the terminal for both
// input and output.
// prompt and prompt are the prompts to display.
func New(prompt string) *Readline {
	if prompt == "" {
		prompt = "> "
	}
	if prompt == "" {
		prompt = ">> "
	}
	return &Readline{
		isTerminal: true,
		prompt:     prompt,
		rdr:        bufio.NewReader(os.Stdin),
		wrt:        os.Stdout,
	}
}

// NewReader creates a new Readline from an existing reader.
func NewReader(rdr *bufio.Reader) *Readline {
	return &Readline{rdr: rdr}
}

// IsTerminal returns true only if the input source is a terminal
func (r *Readline) IsTerminal() bool {
	return r.isTerminal
}

// ReadLine accepts a line of text from the input.
// If the source is the console, then Readline will print
// the prompt before accepting the input.
// If there are no errors reading the source, then all input
// up to end-of-line (or end-of-input) is copied
// into a slice of bytes and returned to the caller.
func (r *Readline) ReadLine() (string, error) {
	if r.isTerminal {
		fmt.Fprintf(r.wrt, "%s", r.prompt)
	}
	line, err := r.rdr.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(line, "\n"), nil
}

// SetPrompt updates the prompt string
func (r *Readline) SetPrompt(prompt string) {
	r.prompt = prompt
}

// Writer is
func (r *Readline) Write(p []byte) (int, error) {
	if r.isTerminal {
		return r.wrt.Write(p)
	}

	return 0, nil
}
