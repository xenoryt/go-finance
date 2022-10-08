package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/DavidGamba/go-getoptions"
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "", 0)
}

func main() {
	logger := Logger()

	opts := getoptions.New()
	filepath := opts.String("csv", "CSV file to load", opts.Alias("f"), opts.Required())

	csvFile, err := os.Open(*filepath)
	if err != nil {
		logger.Fatalln(err)
	}
	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll()
	if err != nil {
		logger.Fatalln(err)
	}
}
