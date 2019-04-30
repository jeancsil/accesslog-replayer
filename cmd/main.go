package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	// "github.com/jeancsil/accesslog-replay/service/apache"
)

func main() {
	start := time.Now()

	fmt.Printf("Starting replaying %s\n\n", "access.log")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// apache.Parse(line)
		// _, err := Parse(scanner.Text())

		// if err != nil {
		// log.Printf("error should not be occured but: %s", err.Error())
		// }
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Printf("\nRan in %s\n", time.Since(start).String())
}
