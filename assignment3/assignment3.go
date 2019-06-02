package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	files := categorizeFiles()
	filePath := "./files/"
	println := fmt.Println
	printf := fmt.Printf
	executionTimeSeq := splitFilesSequentially(filePath, files)
	println("Sequential copying took", executionTimeSeq)
	executionTimeConc := splitFilesConcurrently(filePath, files)
	println("Concurrent copying took", executionTimeConc)

	if executionTimeConc < executionTimeSeq {
		printf("Concurrent method is %f%% faster", float64(100)*(float64(executionTimeSeq)-float64(executionTimeConc))/float64(executionTimeConc))
	} else {
		printf("Sequential method is %f%% faster", float64(100)*(float64(executionTimeConc)-float64(executionTimeSeq))/float64(executionTimeSeq))
	}
}

// Concurrent splitter for files folder
// returns execution time of itself with time.Duration type
func splitFilesConcurrently(filePath string, files map[string][]string) time.Duration {
	println := fmt.Println
	println("Concurrent copy starts")
	startTime := time.Now()
	jpgDone := make(chan string)
	pngDone := make(chan string)
	mp3Done := make(chan string)
	txtDone := make(chan string)
	go splitFilesJPGGO(filePath, files, "./tidy/concurrent/JPG/", jpgDone)
	go splitFilesPNGGO(filePath, files, "./tidy/concurrent/PNG/", pngDone)
	go splitFilesMP3GO(filePath, files, "./tidy/concurrent/MP3/", mp3Done)
	go splitFilesTXTGO(filePath, files, "./tidy/concurrent/TXT/", txtDone)
	<-jpgDone
	<-pngDone
	<-mp3Done
	<-txtDone
	return time.Since(startTime)
}

// Sequential splitter for files folder
// returns execution time of itself with time.Duration type
func splitFilesSequentially(filePath string, files map[string][]string) time.Duration {
	println := fmt.Println
	println("Sequential copy starts")
	startTime := time.Now()

	splitFilesJPG(filePath, files, "./tidy/sequential/JPG/")
	splitFilesPNG(filePath, files, "./tidy/sequential/PNG/")
	splitFilesMP3(filePath, files, "./tidy/sequential/MP3/")
	splitFilesTXT(filePath, files, "./tidy/sequential/TXT/")
	return time.Since(startTime)
}

// Split JPG files sequential method
func splitFilesJPG(filePath string, files map[string][]string, destination string) {
	fmt.Println("Copying JPGs to " + destination)
	for i := range files["jpg"] {
		fileName := files["jpg"][i]
		er := Copy(filePath+fileName, destination+fileName+".jpg")
		if er != nil {
			fmt.Println(er)
		}
	}
}

// Split JPG files concurrent method
func splitFilesJPGGO(filePath string, files map[string][]string, destination string, c chan string) {
	fmt.Println("Copying JPGs to " + destination)
	for i := range files["jpg"] {
		fileName := files["jpg"][i]
		er := Copy(filePath+fileName, destination+fileName+".jpg")
		if er != nil {
			fmt.Println(er)
		}
	}
	c <- "JPG Done"
}

// Split PNG files sequential method
func splitFilesPNG(filePath string, files map[string][]string, destination string) {
	fmt.Println("Copying PNGs to " + destination)
	for i := range files["png"] {
		fileName := files["png"][i]
		er := Copy(filePath+fileName, destination+fileName+".png")
		if er != nil {
			fmt.Println(er)
		}
	}
}

// Split PNG files concurrent method
func splitFilesPNGGO(filePath string, files map[string][]string, destination string, c chan string) {
	fmt.Println("Copying PNGs to " + destination)
	for i := range files["png"] {
		fileName := files["png"][i]
		er := Copy(filePath+fileName, destination+fileName+".png")
		if er != nil {
			fmt.Println(er)
		}
	}
	c <- "PNG Done"
}

// Split MP3 files sequential method
func splitFilesMP3(filePath string, files map[string][]string, destination string) {
	fmt.Println("Copying MP3s to " + destination)
	for i := range files["mp3"] {
		fileName := files["mp3"][i]
		er := Copy(filePath+fileName, destination+fileName+".mp3")
		if er != nil {
			fmt.Println(er)
		}
	}
}

// Split MP3 files concurrent method
func splitFilesMP3GO(filePath string, files map[string][]string, destination string, c chan string) {
	fmt.Println("Copying MP3s to " + destination)
	for i := range files["mp3"] {
		fileName := files["mp3"][i]
		er := Copy(filePath+fileName, destination+fileName+".mp3")
		if er != nil {
			fmt.Println(er)
		}
	}
	c <- "MP3 Done"
}

// Split TXT files sequential method
func splitFilesTXT(filePath string, files map[string][]string, destination string) {
	fmt.Println("Copying TXTs to " + destination)
	for i := range files["txt"] {
		fileName := files["txt"][i]
		er := Copy(filePath+fileName, destination+fileName+".txt")
		if er != nil {
			fmt.Println(er)
		}
	}
}

// Split TXT files concurrent method
func splitFilesTXTGO(filePath string, files map[string][]string, destination string, c chan string) {
	fmt.Println("Copying TXTs to " + destination)
	for i := range files["txt"] {
		fileName := files["txt"][i]
		er := Copy(filePath+fileName, destination+fileName+".txt")
		if er != nil {
			fmt.Println(er)
		}
	}
	c <- "TXT Done"
}

// Simple copy file function
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// returns a map of lists of categorized files according to
// their mime types
func categorizeFiles() map[string][]string {
	files, _ := getDirList("./files/")
	categorized := make(map[string][]string)
	for _, file := range files {
		a := "./files/" + file
		_, ext, _ := mimetype.DetectFile(a)
		if ext == "jpg" {
			categorized["jpg"] = append(categorized["jpg"], file)
		} else if ext == "png" {
			categorized["png"] = append(categorized["png"], file)
		} else if ext == "mp3" {
			categorized["mp3"] = append(categorized["mp3"], file)
		} else if ext == "txt" {
			categorized["txt"] = append(categorized["txt"], file)
		}
	}
	return categorized
}

// Same directory listing function from assignment0
func getDirList(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
