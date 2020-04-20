/*
*	@author: mojtaba.eskandari@waziup.org March 9th 2020
*	@Wazigate Sample APP (Sensors on board)
 */
package main

import (

	"log"
	"net"

	"time"
	"net/http"
	// "encoding/json"
	"strings"

	"os"
	"os/exec"

	routing "github.com/julienschmidt/httprouter"
)

var router = routing.New()

// Please do NOT change this line if you do not know what you are doing
const sockAddr = "/root/app/proxy.sock"

func init() {

	router.GET("/", HomeLink)
	router.GET("/ui/*file_path", UI)
	router.GET("/time", GetTime)
}

/*-------------------------*/

func main() {

	log.Printf("Initializing...")

	if err := os.RemoveAll(sockAddr); err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: router,
	}

	l, e := net.Listen("unix", sockAddr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("Serving... on socket: [%v]", sockAddr)
	server.Serve(l)
}

/*-------------------------*/

// HomeLink serves a hellow world to make sure it works
func HomeLink(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	resp.Write([]byte("Salam Goloooo, It works!"))
}

/*-------------------------*/

// UI serves the static ui
func UI(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	filePath := "." + req.URL.Path
	if filePath == "./" {
		filePath += "index.html"
	}

	http.ServeFile(resp, req, filePath)
}

/*-------------------------*/

// GetTime serves the API /time
func GetTime(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	currentTime := time.Now()
	out := currentTime.Format("2006-01-02 15:04:05 Monday")
	
	zone, _ := currentTime.Zone()
	out += " [ "+ zone +" ]"

	resp.Write([]byte(out ))
}

/*-------------------------*/

// exeCmd executes shell commands
func exeCmd(cmd string, withLogs bool) string {

	if withLogs {
		log.Printf("[Info  ] executing [ %s ] ", cmd)
	}

	exe := exec.Command("sh", "-c", cmd)
	stdout, err := exe.Output()

	if err != nil {
		if withLogs {
			log.Printf("[Err   ] executing [ %s ] command. \n\tError: [ %s ]", cmd, err.Error())

		}
		return ""
	}
	return strings.Trim(string(stdout), " \n\t\r")
}

/*-------------------------*/

// go mod init && go mod tidy
// go build -o start . && ./start