package main

import (
	"fmt"
	"net/http"
        "net/http/httputil"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")

	hostname, error := os.Hostname()

	if error != nil {
     panic(error)
  }


	if len(response) == 0 {
		response = fmt.Sprintf("Hello OpenShift! from %s", hostname)
	}

  requestDump, err := httputil.DumpRequest(r, true)
  if err != nil {
     fmt.Println(err)
  }
  fmt.Println(string(requestDump))

	fmt.Fprintln(w, response)
	fmt.Println("Servicing request.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}


func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "8888"
	}
	go listenAndServe(port)

	select {}
}
