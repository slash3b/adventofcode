package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

// note: try to do this with goroutines ?

func main () {
	key := "yzbqklnj"

	h := md5.New()

	counter := 0
	for {
		att := key + strconv.Itoa(counter)
		_, err := io.WriteString(h, att)
		if err != nil {
			panic(err)
		}

		res := fmt.Sprintf("%x", h.Sum(nil))

		if res[:5] =="00000" {
			fmt.Println(res)
			fmt.Println(counter)
			break
		}

		if res[:6] =="000000" {
			fmt.Println(res)
			fmt.Println(counter)
			break
		}

		counter++
		h.Reset()
	}
}
