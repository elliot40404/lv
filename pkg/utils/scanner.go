package utils

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"github.com/acarl005/stripansi"
)

func ReadPipe(logChannel chan<- string, exitChannel chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = stripansi.Strip(line)
		logChannel <- line
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	exitChannel <- true
}
