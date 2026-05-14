package main

import (
	"bufio"
	"log"
	"os"

	wd "github.com/takanoriyanagitani/go-eword2decoded2047"
)

func sub() error {
	var dec func(string) (string, error) = wd.DecodeNewDefault()

	var bwtr *bufio.Writer = bufio.NewWriter(os.Stdout)
	var scn *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		var line string = scn.Text()
		decoded, err := dec(line)
		if nil != err {
			return err
		}

		_, err = bwtr.WriteString(decoded)
		if nil != err {
			return err
		}

		err = bwtr.WriteByte('\n')
		if nil != err {
			return err
		}
	}

	return bwtr.Flush()
}

func main() {
	err := sub()
	if nil != err {
		log.Printf("%v\n", err)
	}
}
