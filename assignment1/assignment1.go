package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var tpl *template.Template

func init() {
	cwd := getWd()
	tpl = template.Must(template.ParseGlob(cwd + "/template/index.html"))
}

func main() {
	if jailCheck() {
		fmt.Println("Basic test shows that we're not in a jail")
		fmt.Println("Program exiting")
		os.Exit(1)
	} else {
		fmt.Println("Looks like we're in a jail")
		go startWeb()
		// delay needed for serverIP check
		time.Sleep(3 * time.Second)
		DetectServerIP()
		hang()
	}
}
func hang() {
	select {}
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// this test can be bypassed by executing this binary in
// the root folder of the file system "/"
func jailCheck() bool {
	if getWd() != "/" {
		return true
	} else {
		return false
	}
}

func getWd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd
}

func startWeb() {
	http.HandleFunc("/", index)
	http.Handle("/statics/", http.StripPrefix(strings.TrimRight("/statics/", "/"), http.FileServer(http.Dir("template/images"))))
	fmt.Println("WEB SERVER STARTING")
	http.ListenAndServe(":8080", nil)
}

// Prints listening addresses
func DetectServerIP() {
	for IP := range getIP() {
		resp, err := http.Get("http://" + getIP()[IP] + ":8080/")
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Println("Serving on: http://" + getIP()[IP] + ":8080")
		} else {
			fmt.Println("No listening address found!")
		}
	}
}

// get local IP v4 addresses from network interfaces
func getIP() []string {
	internetProtocol := []string{}
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				internetProtocol = append(internetProtocol, ipnet.IP.String())
			}
		}
	}
	return internetProtocol
}
