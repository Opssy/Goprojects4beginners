package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "sync"
)


type Download struct{
	Url string
	DownloadPath string
	FileName string
}

func main(){
	downloads := []Download{
	   { Url: "https://jsonplaceholder.typicode.com/posts/1", DownloadPath: "D:\\", FileName: "file1.json",},
	   { Url: "https://jsonplaceholder.typicode.com/posts/2", DownloadPath: "D:\\", FileName: "file2.json",},
	   { Url: "https://jsonplaceholder.typicode.com/posts/3", DownloadPath: "D:\\", FileName: "file3.json",},
	}
	
	var wg sync.WaitGroup
	 resultfromch := make(chan string, len(downloads)) //channel result
	 errorfromch  := make(chan error, len(downloads)) //channel error

	 for _, download := range downloads {
		wg.Add(1)
		
		go func(download Download) {

			defer wg.Done()
		
		err := DownloadFile(download.Url, download.DownloadPath, download.FileName)
		if err != nil {
			errorfromch <- err
		} else {
			resultfromch <- "File " + download.FileName + " downloaded successfully"
		}
	 }(download)
	}
	go func() {
	wg.Wait()
	close(resultfromch)
	close(errorfromch)
    }()
	   // Print results and errors
    for result := range resultfromch {
        fmt.Println(result)
    }

    for err := range errorfromch {
        fmt.Println("Error:", err)
    }
}

func DownloadFile(url, downloadPath, FileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	 return os.WriteFile(FileName, data, 0644)
}