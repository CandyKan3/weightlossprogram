package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Weight struct {
	Weight []int `json:"weight"`
}

func main() {
	readweight()
	readingjson()
}
func discardBuffer(r *bufio.Reader) {
	r.Discard(r.Buffered())
}
func readingjson() {
	weightarr := readJSON("weight.json")
	fmt.Print("Weights")
	var weightavg float64
	for _, weight := range weightarr {
		var x float64 = float64(weight)
		weightavg += x
	}
	weightavg = weightavg / float64(len(weightarr))
	fmt.Printf("Average weight %g", weightavg)
}
func readJSON(fileName string) []int {
	jsonFile, err := os.Open("weight.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened weight.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//var result map[string]interface{}
	var weights Weight
	json.Unmarshal(byteValue, &weights)
	fmt.Println(weights.Weight[0])
	return weights.Weight

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
