package prices

import (
	"fmt"

	"example.com/data-structures/conversion"
	"example.com/data-structures/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"` // it wil be ignored
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err :=job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringToFoat(lines)
	
	if(err!=nil) {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := (1+job.TaxRate)*price	
		result[fmt.Sprintf("%2.f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	err := job.IOManager.WriteJson(job)
	if err != nil {
		fmt.Println(err)
	}
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
    return &TaxIncludedPriceJob{
        IOManager:   iom,
        InputPrices: []float64{10, 20, 30},
        TaxRate:     taxRate,
    }
}

