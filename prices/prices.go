package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// receiver method to read data from file
func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("could not open file!!")
		fmt.Println(err)
		return
	}
	// bufio is used for reading line by line on a file
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() { // keeps scanning till there is no more content to scan
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err() // find out if an error occured earlier
	if err != nil {
		fmt.Println("reading the file content failed!!")
		fmt.Println(err)
		file.Close() // close the file
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64) // converts text to other value types

		if err != nil {
			fmt.Println("converting price to float failed!!")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
}

// adding a receiver method
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

// constructor
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
