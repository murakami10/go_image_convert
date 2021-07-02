package convert

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func Convert(file File, from, to string) {

	if file.Ext != from || !Checkext(file.Ext) {
		return
	}

	f, err := os.Open(file.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	name := strings.Split(file.Name, ".")[0] + "." + to

	fso, err := os.Create(name)
	if err != nil {
		log.Fatal("create", err)
	}
	defer fso.Close()

	switch to {
	case "jpeg", "jpg":
		jpeg.Encode(fso, img, nil)
		os.Remove(file.Name)
	case "png":
		png.Encode(fso, img)
		os.Remove(file.Name)
	case "gif":
		gif.Encode(fso, img, nil)
		os.Remove(file.Name)
	default:
	}

	return
}
