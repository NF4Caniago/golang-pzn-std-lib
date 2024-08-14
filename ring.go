package main
import (
	"fmt"
	"container/ring"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(6)
	// data.Value = "Value 1"
	
	// data = data.Next()
	// data.Value = "Value 2"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any){
		fmt.Println(value)
	})
}