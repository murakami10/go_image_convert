package main

import (
	"flag"
	"kadai1/convert"
	"log"
)

func main() {

	from := flag.String("from", "jpg", "extension")
	to := flag.String("to", "png", "extension")
	dirname := flag.String("dir", "", "dir")

	flag.Parse()

	if *dirname == "" {
		log.Fatal("dirname is empty")
	}

	if !convert.Checkext(*from) || !convert.Checkext(*to) {
		log.Fatal("from:", *from, " to:", *to)
	}

	files := convert.Tofile(convert.Dirsearch(*dirname))

	for _, filename := range files {
		convert.Convert(filename, *from, *to)
	}
}
