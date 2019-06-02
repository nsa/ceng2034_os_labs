# Assignment 3 - Final

[![forthebadge](https://forthebadge.com/images/badges/pretty-risque.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com)

## Task

You have a chunk of files (mixed png, jpeg, txt). But these files have no extension.
- Categorize the files according to their types. (copy or move each of them to different folders such as jpegFolder, pngFolder, txtFolder...)
- Use concurrency methodology to do this job more faster.
- Do a benchmark to show the difference between single thread and multi thread.

There are 4 types of files in files directory:
- file1 jpg
- file2 mp3
- file3 png
- file4 ASCI txt
(other 140 files are mixed of them. You can look over `copyScript.py` for how i do that.)

## Information

**`assignment3.go`** reads chunked files from the folder called ***files*** and categorizes them according to their MIME type then splits the files to another folder called ***tidy***. 

Its **sequential** part starts splitting files with the following order JPG, PNG, MP3, TXT. When sequential part finishes it stores the execution time of this part to a variable for after use.

After the sequential part, it starts splitting files **concurrently** and stores its execution time to another variable.

After all splitting work finishes the program prints which method is faster and by how much(%).

## Dependencies

- Go: go1.11.5
- Docker: 18.09.6-ce
- Mimetype(github.com/gabriel-vasile/mimetype)

## Instructions

* Clone this repository to your computer
    - `git clone https://github.com/nsa/ceng2034_os_labs`

* Change working directory into the folder **assignment3**
    - `cd assignment3`

### Run with Docker(recommended)

* Build Docker image from the Dockerfile
    - `docker build -t os_assignment3 .`

* Run it in a container
    - `docker run -it --rm os_assignment3`

### Run without Docker(optional)
* Install dependency **Mimetype**
    - `go get github.com/gabriel-vasile/mimetype`

* Run **`tidyTidy.sh`** script to create necessary folders.
 	- `bash tidyTidy.sh`

* Run **`assignment3.go`** to categorize files
    - `go run assignment3.go`

* (Optional) Build **`assignment3.go`**
    - `go build assignment3.go` 

* (Optional) Execute **`assignment3`** binary
    - `./assignment3`


## Expected result

```
Sequential copy starts
Copying JPGs to ./tidy/sequential/JPG/
Copying PNGs to ./tidy/sequential/PNG/
Copying MP3s to ./tidy/sequential/MP3/
Copying TXTs to ./tidy/sequential/TXT/
Sequential copying took 21.841689ms
Concurrent copy starts
Copying TXTs to ./tidy/concurrent/TXT/
Copying PNGs to ./tidy/concurrent/PNG/
Copying JPGs to ./tidy/concurrent/JPG/
Copying MP3s to ./tidy/concurrent/MP3/
Concurrent copying took 12.162092ms
Concurrent method is 79.588257% faster
```