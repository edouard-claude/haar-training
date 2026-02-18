package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marianina8/haar-training/images"
	"github.com/marianina8/haar-training/samples"
	"github.com/marianina8/haar-training/training"
)

func main() {

	// 1. DOWNLOAD NEGATIVE BACKGROUND FILES

	links := []string{
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n12102133",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n09436708",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n12992868",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n07942152",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n02913152",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n04105893",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n03089879",
		"http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n10529231",
	}
	images.Get(links, "negatives", true, 1, 4000, 200, 200)

	// 2. GENERATE BG.TXT FILE FROM DOWNLOADED NEGATIVE (BACKGROUND) FILES

	files, err := os.ReadDir("negatives")
	if err != nil {
		fmt.Println("err reading dir:", err)
	}
	var data string
	for _, file := range files {
		data += "negatives/" + file.Name() + "\n"
	}
	err = os.WriteFile("bg.txt", []byte(data), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 3. CREATE POSITIVE SAMPLE VECTOR FILE
	createSampleCmdOptions := "-maxxangle 0.5 -maxyangle 0.5 -maxzangle 0.5"
	samples.CreateSamples("adidas.png", "bg.txt", 1950, createSampleCmdOptions)

	// 4. TRAIN HAAR CASCADE FILE
	training.HaarCascade("data", 1800, 900, 20)
}
