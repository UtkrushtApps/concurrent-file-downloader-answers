// concurrent-file-downloader/main.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DownloadResult struct {
	FileName string
	Status   string
}

func simulateDownload(fileName string, ch chan<- DownloadResult, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate random download time between 2-3 seconds
	duration := time.Duration(rand.Intn(1000)+2000) * time.Millisecond
	time.Sleep(duration)
	ch <- DownloadResult{
		FileName: fileName,
		Status:   "completed",
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt"}
	ch := make(chan DownloadResult) // unbuffered channel
	var wg sync.WaitGroup

	for _, fname := range files {
		wg.Add(1)
		go simulateDownload(fname, ch, &wg)
	}

	// Goroutine to close channel when all downloads are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive and print completion messages as they arrive
	for res := range ch {
		fmt.Printf("%s: %s\n", res.FileName, res.Status)
	}

	// All done, exit cleanly
}