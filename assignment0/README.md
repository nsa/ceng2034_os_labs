# Assignment 0 - import os

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/uses-git.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com)

## Task

* Go to home directory and create a folder name "os_lab_0".
* Change working directory to new created folder.
* Create new 2 .txt file and 1 .py file in the folder. (!must be empty files)
* Print all files last modified date in the folder.
* Find only txt files in the folder.
* **!Be sure your code does not require sudo privilege on any unix-like os.**

## Information

* Written in Golang **1.11.5**
* Shipped with Docker **18.09.2-ce**

## Instructions

* Clone this repository to your computer
    - `git clone https://github.com/nsa/ceng2034_os_labs`

* Change working directory into the folder **assignment0**
    - `cd assignment0`

### Run with Docker(recommended)

* Build Docker image from the Dockerfile
    - `docker build -t os_assignment0 .`

* Run it in a container
    - `docker run os_assignment0`

### Run without Docker

* Run **`assignment0.go`**
    - `go run assignment0.go`

    or

* Compile and Run **`assignment0.go`**
    - `go build assignment0 .`
    - `./assignment0`

## Expected output

```
Home:  /root
Pwd:  /some_random/working_dir
Creating new folder into Home: os_lab_0
Changing working directory to: os_lab_0
New Pwd:  /root/os_lab_0
Creating empty files
Last mod of textfile_one.txt is: 2019-02-28 21:53:45.111918767 +0000 UTC
Last mod of some_random_empty_python_file.py is: 2019-02-28 21:53:45.111918767 +0000 UTC
Last mod of textfile_two.txt is: 2019-02-28 21:53:45.111918767 +0000 UTC
Txt file: textfile_one.txt
Txt file: textfile_two.txt
```