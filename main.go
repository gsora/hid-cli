package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gsora/hid-compiler/compiler"
)

const (
	hidPath = "/dev/hidg0"
)

func main() {
	if _, err := os.Stat(hidPath); os.IsNotExist(err) {
		log.Fatal("cannot find hid interface on", hidPath)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("hid-cli> ")
		input, _ := reader.ReadString('\n')
		compiledInput := compiler.Compile(input)
		compiledInput = compiledInput[:len(compiledInput)-1]
		err := writeToHID(compiledInput)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeToHID(s string) error {
	f, err := os.OpenFile(hidPath, os.O_WRONLY, os.ModeCharDevice)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(s))
	return err
}
