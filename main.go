package main

import (
	"fmt"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	var fileName string
	var height, width int
	if len(os.Args) == 4 {
		fileName = os.Args[1]
		width, _ = strconv.Atoi(os.Args[3])
		height, _ = strconv.Atoi(os.Args[2])
	} else {
		fmt.Print("input text: ")
		_, err := fmt.Scanln(&fileName, &width, &height)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

	imageFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening input image:", err)
		return
	}
	inputImage, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println("Error decoding input image:", err)
		return
	}

	if height > 0 || width > 0 {
		inputImage = resize.Resize(uint(width), uint(height), inputImage, resize.Lanczos3)
	}

	widthStr := strconv.Itoa(inputImage.Bounds().Dx())
	heightStr := strconv.Itoa(inputImage.Bounds().Dy())
	for len(widthStr) < 3 {
		widthStr = "0" + widthStr
	}
	for len(heightStr) < 3 {
		heightStr = "0" + heightStr
	}

	imageOut := widthStr + heightStr

	for y := inputImage.Bounds().Dy(); y > 0; y-- {
		for x := 0; x < inputImage.Bounds().Dx(); x++ {
			r, g, b, _ := inputImage.At(x, y).RGBA()
			rStr := strconv.Itoa(int(r) / 256)
			gStr := strconv.Itoa(int(g) / 256)
			bStr := strconv.Itoa(int(b) / 256)
			for len(rStr) < 3 {
				rStr = "0" + rStr
			}
			for len(gStr) < 3 {
				gStr = "0" + gStr
			}
			for len(bStr) < 3 {
				bStr = "0" + bStr
			}
			imageOut += rStr + gStr + bStr
		}
	}
	fmt.Println(imageOut)
}
