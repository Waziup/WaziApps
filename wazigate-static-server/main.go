/*
*	@author: mojtaba.eskandari@waziup.org March 9th 2020
*	@Wazigate Sample APP (Sensors on board)
 */
package main

import (

	"log"
	"net"

	// "time"
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
	router.GET("/*file_path", UI)
}

/*-------------------------*/

func main() {

	log.Printf("Initializing...")

	// copyTheUIDirectory();

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
/*
func copyTheUIDirectory(){

	log.Printf("Checking if the UI directory exist...");

	empty, err := isEmpty( "/root/ui");
	if( empty || err != nil){

		log.Printf("Couldn't find the UI directory, copying the default UI...");

		exeCmd( "cp "

	}else{
		log.Printf("The UI directory exists.");
	}

}


/*-------------------------*/
/*
func isEmpty(name string) (bool, error) {
    f, err := os.Open(name)
    if err != nil {
        return false, err
    }
    defer f.Close()

    _, err = f.Readdirnames(1) // Or f.Readdir(1)
    if err == io.EOF {
        return true, nil
    }
    return false, err // Either not empty or error, suits both cases
}

/*-------------------------*/

// UI serves the static ui
func UI(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	filePath := req.URL.Path
	if filePath == "/" {
		filePath += "index.html"
	}

	log.Printf("[WWW  ] >> %s %q", req.Method, req.URL)

	http.ServeFile(resp, req, "/root/ui/" + filePath)
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