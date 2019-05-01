package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jeancsil/accesslog-replayer/internals/contentparser"
	"github.com/jeancsil/accesslog-replayer/internals/service"
)

func main() {
	start := time.Now()

	fmt.Printf("Starting replaying %s\n\n", os.Stdin.Name())

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		// apache := contentparser.Apache{}
		nginx := contentparser.Nginx{}
		replay := service.Replay{
			ContentServer: nginx,
		}
		replay.Run(line)

		// _, err := Parse(scanner.Text())

		// if err != nil {
		// log.Printf("error should not be occured but: %s", err.Error())
		// }
	}

	// if err := scanner.Err(); err != nil {
	// 	log.Println(err)
	// }

	fmt.Printf("\nRan in %s\n", time.Since(start).String())
}
