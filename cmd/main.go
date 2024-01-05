package main

import (
	"fmt"
	"os"

	ctm "github.com/harakeishi/ctm"
)

func main() {
	md := ctm.NewMarkDownDataSet(os.Stdin)
	fmt.Printf("%s", md)
}
