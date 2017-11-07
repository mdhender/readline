// Package readline implements a very simple console reader.
package readline

import (
	"bufio"
	"io"
	"os"
)

// ReadWriter should satisfy a buffered Reader interface and unbuffered Writer interface.
type ReadWriter struct {
	prompt        []byte
	*bufio.Reader // console input
	io.Writer     // console output
}

// NewReadWriter creates a new reader by opening the console for both input and output.
// Prompt is the prompt to display before accepting input.
func NewReadWriter(prompt string) *ReadWriter {
	rw := &ReadWriter{
		prompt: []byte("> "),
		Reader: bufio.NewReader(os.Stdin),
		Writer: os.Stdout,
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
		tmp, isPrefix, err := rw.Reader.ReadLine()
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

// ReadString prints a prompt, then accepts a line of text from the console.
// The implementation is differs from Reader.ReadString in that end-of-line is treated wonkily.
// If the delimiter is '\n', then the end-of-line characters are not included in the string returned.
// (In other words, both '\r' and '\n' are trimmed from the end of the string).
// Otherwise, all input up to (and including) the delimiter is returned.
func (rw *ReadWriter) ReadString(delim byte) (string, error) {
	if delim != '\n' {
		rw.Prompt()
		return rw.Reader.ReadString(delim)
	}
	line, _, err := rw.ReadLine()
	return string(line), err
}

// ReadToEOL prints a prompt, then accepts a line of text from the console.
func (rw *ReadWriter) ReadToEOL() (string, error) {
	bytes, _, err := rw.ReadLine()
	return string(bytes), err
}

// SetPrompt updates the prompt string
func (rw *ReadWriter) SetPrompt(prompt string) {
	rw.prompt = []byte(prompt)
}
