package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/frozzare/go-util/httputil"
)

type response struct {
	Answer []struct {
		Name string
		Type int
		TTL  int
		Data string
	}
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
			if len(res.Answer) == 0 || res.Answer[0].Name == "." {
				fmt.Println("No response from " + os.Args[1])
			} else {
				for _, answer := range res.Answer {
					data := answer.Data

					if strings.HasSuffix(data, ".") {
						data = data[:len(data)-1]
					}

					fmt.Println(fmt.Sprintf("Using %s with TTL %d", data, answer.TTL))
				}
			}
		}
	}
}
