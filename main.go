package main

import (
	"fmt"
	"os"

	"github.com/frozzare/go-util/httputil"
)

type response struct {
	Comment string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing domain argument")
	} else {
		var res response
		err := httputil.GetJSON("https://dns.google.com/resolve?name="+os.Args[1], &res)

		if err != nil {
			fmt.Println(err)
		} else {
			if res.Comment == "" {
				fmt.Println("No response from " + os.Args[1])
			} else {
				fmt.Println(res.Comment)
			}
		}
	}
}
