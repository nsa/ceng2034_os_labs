package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	home := os.Getenv("HOME")
	fmt.Println("Home: ", home)

	fmt.Println("Pwd: ", get_wd())
	os.Chdir(home)
	folder_name := "os_lab_0"
	fmt.Println("Creating new folder:", folder_name)
	os.Mkdir(folder_name, 0775)
	fmt.Println("Changing working directory to:", folder_name)
	os.Chdir(folder_name)
	fmt.Println("New Pwd: ", get_wd())

	create_empty_files()
	get_last_mod()
	get_text_files()
}

func get_wd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd
}

func create_empty_files() {
	os.Create("textfile_one.txt")
	os.Create("textfile_two.txt")
	os.Create("some_random_empty_python_file.py")
}

func get_dir_list(root string) ([]string, error) {
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

func get_last_mod() {
	dir_list, _ := get_dir_list(get_wd())
	for _, file := range dir_list {
		f, _ := os.Stat(file)
		fmt.Println("Last mod of", file, "is:", f.ModTime())
	}
}

func get_text_files() {
	dir_list, _ := get_dir_list(get_wd())
	for _, file := range dir_list {
		extension := filepath.Ext(file)
		if extension == ".txt" {
			fmt.Println("Txt file:", file)
		}
	}
}
