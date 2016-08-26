package main

import "flag"
import "fmt"
import "net/http"
import "log"
import "io"
import "os"

func main() {
	flag.Parse()
	fmt.Println(flag.Args())

	resp, err := http.Get("http://<version endpoint>")

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println()
}
