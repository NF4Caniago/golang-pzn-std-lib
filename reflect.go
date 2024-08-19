package main
import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string
	Address string
	Email string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name , "with type", valueField.Type)
		fmt.Println(StructField.Tag.Get("required"))
		fmt.Println(StructField.Tag.Get("max"))
	}
}

func main() {
	readField(Sample{"Eko"})
	readField(Person{"Eko", "", ""})
}