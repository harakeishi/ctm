package ctm

import (
	"fmt"
	"io"
)

type MarkDownDataSet struct {
	input io.Reader
}

func NewMarkDownDataSet(in io.Reader) *MarkDownDataSet {
	fmt.Printf("%s", in)
	return &MarkDownDataSet{in}
}
