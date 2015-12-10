package godl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadFromURL(url string, fileName string) error {
	if fileName == "" {
		tokens := strings.Split(url, "/")
		fileName = tokens[len(tokens)-1]
	}
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error while creating", fileName)
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error while downloading", url)
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return fmt.Errorf("Error while downloading", url)
	}

	fmt.Println(n, "bytes downloaded to", fileName)
	return nil
}

func DownloadFromURLParallel(url string, fileName string, done chan bool) {
	err := DownloadFromURL(url, fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	done <- true
}
