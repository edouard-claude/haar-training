package images

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
)

// Get downloads images from ImageNet URLs, resizes and optionally converts them to grayscale.
func Get(links []string, folderName string, grayscale bool, start, limit, height, width int) {
	fmt.Println("inside get")
	client := http.DefaultClient

	picNum := start
	for _, imageLink := range links {
		fmt.Println("link:", imageLink)
		req, err := http.NewRequest(http.MethodGet, imageLink, nil)
		if err != nil {
			continue
		}
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		// create folder if does not exist
		mode := int64(0777)
		if _, err := os.Stat(folderName); os.IsNotExist(err) {
			os.MkdirAll(folderName, os.FileMode(mode))
		}
		// scan through each line of the response body
		scanner := bufio.NewScanner(resp.Body)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			fmt.Println("scanning line:", scanner.Text())
			if picNum > limit {
				resp.Body.Close()
				return
			}
			// get the image associated with the link
			imgResp, e := http.Get(scanner.Text())
			if e != nil {
				continue
			}
			if imgResp.StatusCode != http.StatusOK {
				imgResp.Body.Close()
				continue
			}
			// open a file for writing
			filePath := filepath.Join(folderName, strconv.Itoa(picNum)+".jpg")
			file, err := os.Create(filePath)
			if err != nil {
				imgResp.Body.Close()
				fmt.Println("error creating:", err)
				continue
			}
			// Use io.Copy to just dump the response body to the file. This supports huge files
			n, err := io.Copy(file, imgResp.Body)
			imgResp.Body.Close()
			file.Close()
			if err != nil || n < 3000 {
				_ = os.Remove(filePath)
				continue
			}
			// open the file for image manipulation
			srcImg, err := imaging.Open(filePath)
			if srcImg == nil || err != nil {
				fmt.Println("error opening:", err)
				continue
			}
			// resize image
			if height > 0 || width > 0 {
				srcImg = imaging.Resize(srcImg, width, height, imaging.Lanczos)
			}
			// convert image to grayscale
			if grayscale {
				srcImg = imaging.Grayscale(srcImg)
			}
			err = imaging.Save(srcImg, filePath)
			if err != nil {
				fmt.Println("error saving:", err)
				continue
			}
			picNum++
		}
		resp.Body.Close()
	}
}
