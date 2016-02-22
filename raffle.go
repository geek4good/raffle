package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	filename, err := filenameFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file := os.Stdin
	if filename != "" {
		if file, err = os.Open(filename); err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	speakers, err := linesAsArray(file)
	if err != nil {
		log.Fatal(err)
	}

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("")
	fmt.Println("And the winner is...")
	fmt.Println("")
	time.Sleep(1000 * time.Millisecond)

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	fmt.Println(speakers[random.Intn(len(speakers))])
}

func filenameFromCommandLine() (filename string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		t := time.Now()
		err = fmt.Errorf("Usage: %s [speakers_%d-%02d-%02d.txt]", filepath.Base(os.Args[0]), t.Year(), t.Month(), t.Day())
		return "", err
	}
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	return filename, nil
}

func linesAsArray(file io.Reader) (speakers []string, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			speakers = append(speakers, scanner.Text())
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return speakers, nil
}
