package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	input, err := io.ReadAll(ch.Reader)
	if err != nil {
		return err
	}

	expr := string(input)

	result, err := EvalPostfix(expr)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(ch.Writer, result)
	return err
}
