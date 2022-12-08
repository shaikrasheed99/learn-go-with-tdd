package csvtesting

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadCSV(writer io.Writer) {
	file, error := os.Open("books.csv")

	if error != nil {
		fmt.Fprintln(writer, error.Error())
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, error := csvReader.ReadAll()

	if error != nil {
		fmt.Fprintln(writer, error.Error())
	}

	fmt.Fprintln(writer, data[1])
}
