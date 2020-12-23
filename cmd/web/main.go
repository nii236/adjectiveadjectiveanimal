package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	aaa "github.com/nii236/adjectiveadjectiveanimal"
)

var port int
var server bool

func init() {
	flag.IntVar(&port, "port", 8080, "what port to serve on")
	flag.BoolVar(&server, "server", false, "run server")
}

func main() {
	flag.Parse()
	if server {
		serve()
		return
	}
	args := os.Args[1:]
	i := 2
	if len(args) > 0 {
		var err error
		i, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Non number supplied for adjective amount. Exiting.")
			os.Exit(-1)
		}
	}
	result := strings.Join(aaa.Generate(i, &aaa.Options{}), "-")
	fmt.Println(result)
}

func serve() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sep := r.URL.Query().Get("sep")

		if sep == "" {
			sep = "-"
		}
		numStr := r.URL.Query().Get("num")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			num = 2
		}
		words := aaa.Generate(num, &aaa.Options{})

		w.Write([]byte(strings.Join(words, sep)))
	})

	portNum := strconv.Itoa(port)
	log.Println("Run server on:", portNum)
	log.Fatalln(http.ListenAndServe(":"+portNum, m))
}
