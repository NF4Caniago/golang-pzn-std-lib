package main
import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Afif", 27},
		{"Eko", 30},
		{"Lala", 22},
		{"Raline", 4},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}