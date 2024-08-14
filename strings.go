package main
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("Eko Kurniawan", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan"))
	fmt.Println(strings.ToUpper("Afif Ilham Caniago"))
	fmt.Println(strings.Trim(" Afif Ilham Caniago  ", " "))
	fmt.Println(strings.ReplaceAll("Afif Ilham Caniago", "Afif", "Eko"))
}