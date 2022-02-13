package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
)

func main() {
	i := 512
	for i > 0 {
		RandomString(28)
		i -= 1
	}
}
func makefile(randstring string) {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	var pathst string
	var pathnd string
	pathst = user.HomeDir
	pathnd = "/Desktop/"
	var path string = pathst + pathnd
	fmt.Println(path)
	flooderfile, err := os.Create(path + "floodedfile" + randstring + ".txt")
	flooderfile.WriteString(":)))))")
	flooderfile.Close()

	fmt.Println("FileCreated")
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	makefile(string(s))
	return string(s)
}
