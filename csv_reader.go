package main
import (
	"fmt"
	"encoding/csv"
	"io"
	"strings"
)

func main() {
	csvString := "Afif, Ilham, Caniago\n" + "Latifah, Humaira, Jelek\n" + "Lailatur, Rahmi, Tengak"
	
		reader := csv.NewReader(strings.NewReader(csvString))

		for{
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			fmt.Println(record)
		}
}