package save

import (
	"fmt"
	"log"

	"crypto-mine-cli/cmd/commands"

	"github.com/gocolly/colly"
)

func saveInCSV(fileName string) error {
	file, writer, err := createCSVFile(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	defer writer.Flush()

	if err := writer.Write(
		[]string{
			"Name",
			"Symbol",
			"Market capacity (USD)",
			"Price (USD)",
			"Volume (USD)",
			"Change (1h)",
			"Change (24h)",
			"Change (7d)",
		},
	); err != nil {
		return fmt.Errorf("error in writer csv header: %v", err)
	}

	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		metrics := commands.GetMetrics(h)

		if err := writer.Write([]string{
			metrics.Name,
			metrics.Symbol,
			metrics.MarketCap,
			metrics.Price,
			metrics.Volume,
			metrics.Change1h,
			metrics.Change24h,
			metrics.Change7d,
		}); err != nil {
			log.Printf("failure to write crypto %s line: %v", metrics.Name, err)
		}
	})

	c.OnError(func(_ *colly.Response, requestErr error) {
		err = fmt.Errorf("Something went wrong: %v", requestErr)
	})

	if err := c.Visit(commands.CoinMarketURL); err != nil {
		return fmt.Errorf("failed in visit %s", commands.CoinMarketURL)
	}

	return nil
}
