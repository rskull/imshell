package main

import (
	"flag"
	"fmt"

	"github.com/rskull/imshell/imshell"
)

func main() {

	const (
		version      = "1.0.0"
		defaultWidth = 50
	)

	var (
		imagePath      string
		width          float64
		text           string
		versionPrinted bool
	)

	flag.StringVar(&imagePath, "i", "", "image file path")
	flag.Float64Var(&width, "w", defaultWidth, "output width")
	flag.StringVar(&text, "t", "", "given text replace")
	flag.BoolVar(&versionPrinted, "v", false, "imshell version")
	flag.Parse()

	if versionPrinted {
		fmt.Println(version)
		return
	}

	if imagePath == "" {
		fmt.Println("image file path or image url")
		return
	}

	buffer, err := imshell.Convert(imagePath, text, width)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(buffer)
}
