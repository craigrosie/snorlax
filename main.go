package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	seconds := flag.Float64("s", 1.0, "seconds to sleep for")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		time.Sleep(time.Duration(*seconds*1000) * time.Millisecond)
	}
}
