// Package readline implements a very simple console reader.
package readline

import (
	"bufio"
	"io"
	"os"
)

// ReadWriter should satisfy a buffered Reader interface and unbuffered Writer interface.
type ReadWriter struct {
	console *bufio.Reader // always buffered
	io.Writer
	prompt []byte
}

// NewReadWriter creates a new reader by opening the console for both input and output.
// Prompt is the prompt to display before accepting input.
func NewReadWriter(prompt string) *ReadWriter {
	rw := &ReadWriter{
		console: bufio.NewReader(os.Stdin),
		Writer:  os.Stdout,
		prompt:  []byte("> "),
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
	return rw.Writer.Write(rw.prompt)
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
	rw.Prompt()
	return rw.console.ReadString(delim)
}

// ReadToEOL prints a prompt, then accepts a line of text from the console.
func (rw *ReadWriter) ReadToEOL() (string, error) {
	rw.Prompt()
	bytes, _, err := rw.ReadLine()
	return string(bytes), err
}

// SetPrompt updates the prompt string
func (rw *ReadWriter) SetPrompt(prompt string) {
	rw.prompt = []byte(prompt)
}
