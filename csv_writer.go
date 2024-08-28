package main
import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"afif", "ilham", "caniago"})
	_ = writer.Write([]string{"latifah", "humaira", "jelek"})
	_ = writer.Write([]string{"lailatur", "rahmi", "tengak"})

	writer.Flush()
}