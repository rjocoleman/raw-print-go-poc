package main

import (
	"fmt"
	"github.com/alexbrainman/printer"
	"os"
)

func main() {
	const zpl = "^XA^CFA,30^FO50,50^FDHello %s^FS^XZ"

	name, err := printer.Default()
	if err != nil {
		handleErr(err)
	}
	fmt.Printf("Default Printer is: %s\n", name)
	fmt.Println("Starting to send print job")

	p, err := printer.Open(name)
	if err != nil {
		handleErr(err)
	}
	defer p.Close()

	err = p.StartDocument("Raw Print Test", "RAW")
	if err != nil {
		handleErr(err)
	}
	defer p.EndDocument()

	err = p.StartPage()
	if err != nil {
		handleErr(err)
	}

	fmt.Fprintf(p, zpl, name)

	err = p.EndPage()
	if err != nil {
		handleErr(err)
	}
	fmt.Println("Finished sending print job")
}

func handleErr(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}
