// Package readline implements a very simple console reader.
package readline

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadWriter should satisfy a buffered Reader interface.
// For convenience, we use "console" for input, and terminal for "output".
type ReadWriter struct {
	console  *bufio.Reader // always buffered
	terminal io.Writer     // defaults to unbuffered
	prompt   []byte
}

// NewReadWriter creates a new reader by opening the console for both input and output.
// Prompt is the prompt to display before accepting input.
func NewReadWriter(prompt string) *ReadWriter {
	rw := &ReadWriter{
		console:  bufio.NewReader(os.Stdin),
		terminal: os.Stdout,
		prompt:   []byte("> "),
	}
	if prompt != "" {
		rw.SetPrompt(prompt)
	}
	return rw
}

// GetPrompt returns the prompt string
func (rw *ReadWriter) GetPrompt() string {
	return string(rw.prompt)
}

// Prompt writes the prompt string to the terminal
func (rw *ReadWriter) Prompt() (n int, err error) {
	if len(rw.prompt) == 0 {
		return 0, nil
	}
	return rw.terminal.Write(rw.prompt)
}

// ReadLine prints a prompt, then accepts a line of text from the console.
// If the line is returned, it will not contain the end-of-line character(s).
// Unlike Reader.ReadLine, this implementation gathers the entire line.
// Even though it returns isPrefix to satisfy the interface, it always returns the entire line or an error.
func (rw *ReadWriter) ReadLine() (line []byte, isPrefix bool, err error) {
	rw.Prompt()
	for {
		tmp, isPrefix, err := rw.console.ReadLine()
		if err != nil {
			if err == io.EOF && line != nil {
				return line, false, nil
			}
			return nil, false, err
		}
		line = append(line, tmp...)
		if !isPrefix {
			break
		}
	}
	return line, isPrefix, err
}

// ReadRune reads a single rune
// need to embed this in

// ReadString prints a prompt, then accepts a line of text from the console.
// If there are no errors reading the source, then all input up to delimiter is returned to the caller.
func (rw *ReadWriter) ReadString(delim byte) (string, error) {
	if len(rw.prompt) > 0 {
		rw.Write(rw.prompt)
	}
	return rw.console.ReadString(delim)
}

// ReadSString prints a prompt, then accepts a line of text from the console.
// If there are no errors reading the source, then all input up to end-of-line (or end-of-input) is returned to the caller.
func (rw *ReadWriter) ReadSString(delim byte) (string, error) {
	if len(rw.prompt) > 0 {
		rw.Write(rw.prompt)
	}
	line, err := rw.console.ReadString('\n')
	if err == nil {
		strings.TrimRight(line, "\r\n")
	}
	return line, err
}

// SetPrompt updates the prompt string
func (rw *ReadWriter) SetPrompt(prompt string) {
	rw.prompt = []byte(prompt)
}

// Write a slice of bytes to the terminal
func (rw *ReadWriter) Write(p []byte) (n int, err error) {
	return rw.terminal.Write(p)
}
