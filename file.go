package main
import (
	"os"
	"io"
	//"fmt"
	"bufio"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666 )
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666 )
	if err != nil {
		return "", err
	}
	defer file.Close()

	var message string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666 )
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message + "\n")
	return nil
}

func main() {
	//createNewFile("newlog.log", "sample of log")

	// file, _ := readFile("newlog.log")
	// fmt.Println(file)

	addToFile("newlog.log", "new log added")
}