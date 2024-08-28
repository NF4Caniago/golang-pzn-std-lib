package main
import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Afif", "Paul", "tomo", "mulyono"}
	values := []int{100, 98, 70, 40}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Eko"))
	fmt.Println(slices.Index(names, "Eko"))
	fmt.Println(slices.Index(names, "Paul"))
}