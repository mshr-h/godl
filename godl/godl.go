package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/mshr-h/godl"
)

func main() {
	parallels := 0
	done := make(chan bool)

	var inputFileName = flag.String("i", "", "list of urls")
	flag.Parse()

	if *inputFileName == "" {
		urlList := os.Args[1:]
		for i := 0; i < len(urlList); i++ {
			url := urlList[i]
			parallels++
			go godl.DownloadFromURLParallel(url, "", done)
		}
	} else {
		fp, err := os.Open(*inputFileName)
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			url := scanner.Text()
			parallels++
			go godl.DownloadFromURLParallel(url, "", done)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	for i := 0; i < parallels; i++ {
		<-done
	}
}
