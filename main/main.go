package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

 

func main(){
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Println(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didnt enter your name")
	}

	return name, nil
}
