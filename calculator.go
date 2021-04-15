package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Weight struct {
	Weight float64
	Date   string
}

func main() {
	weight := readweight()
	readingjson(weight)
}
func discardBuffer(r *bufio.Reader) {
	r.Discard(r.Buffered())
}
func readingjson(weightt float64) {
	weightarr := readJSON("weight.json", weightt)
	//fmt.Print("Weights")
	var weightavg float64
	for _, weight := range weightarr {
		var x float64 = float64(weight.Weight)
		weightavg += x
		fmt.Print(x, ",")
	}
	weightavg = weightavg / float64(len(weightarr))
	fmt.Println()
	fmt.Printf("Average weight %g", weightavg)
	fmt.Println()
	fmt.Printf("You are %g pounds down from your original weight", weightarr[0].Weight-weightarr[len(weightarr)-1].Weight)
}
func readJSON(fileName string, i float64) []Weight {
	jsonFile, err := os.Open("weight.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var weights []Weight

	json.Unmarshal(byteValue, &weights)

	dt := time.Now()
	weights = append(weights, Weight{Weight: i, Date: dt.Format("04-014-2021")})
	file, _ := json.MarshalIndent(weights, "", "")
	_ = ioutil.WriteFile("weight.json", file, 0644)
	return weights

}
func readweight() float64 {
	stdin := bufio.NewReader(os.Stdin)
	var i float64
	_, err := fmt.Fscanln(stdin, &i)
	if err != nil {
		//throw error, loop
		discardBuffer(stdin)
		fmt.Println("Error: Please enter a valid weight")
		readweight()
	} else {
		fmt.Print("read number", i, "from stdin")
		return i
	}
	return i
}
