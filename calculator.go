package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readweight()
}
func discardBuffer(r *bufio.Reader) {
	r.Discard(r.Buffered())
}
func readweight() {
	stdin := bufio.NewReader(os.Stdin)
	var i int
	_, err := fmt.Fscanln(stdin, &i)
	if err != nil {
		//throw error, loop
		discardBuffer(stdin)
		fmt.Println("Error: Please enter a valid weight")
		readweight()
	} else {
		fmt.Print("read number", i, "from stdin")
	}
}
