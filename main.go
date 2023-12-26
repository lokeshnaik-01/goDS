package main

import (
	"example.com/data-structures/prices"
	"example.com/data-structures/filemanager"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

}
