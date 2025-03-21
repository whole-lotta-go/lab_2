package lab2

import "io"

type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	return nil
}
