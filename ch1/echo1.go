//ECHO1 prints out its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = sep + os.Args[i]
		fmt.Printf("%v %s \n", i, s)
	}

}
