package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type LogLevel string

const (
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARNING"
	Error   LogLevel = "ERROR"
	Debug   LogLevel = "DEBUG"
)

func main() {
	duration := flag.Int("duration", 5, "Duration in seconds for log generation")
	interval := flag.Int("interval", 500, "Interval in milliseconds between each log")
	infinite := flag.Bool("infinite", false, "Generate logs infinitely")
	flag.Parse()
	// Parse the input argument as the duration in seconds
	durationSec := *duration
	// Parse the input argument as the interval in milliseconds
	intervalMs := *interval
	// Calculate the total duration in milliseconds
	durationMs := durationSec * 1000
	// Create a ticker that generates logs every intervalMs milliseconds
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	// Start generating logs
	if *infinite {
		for range ticker.C {
			log := GenerateRandomLog()
			fmt.Println(log)
		}
	} else {
		start := time.Now()
		for now := range ticker.C {
			elapsed := now.Sub(start)
			if elapsed > time.Duration(durationMs)*time.Millisecond {
				break
			}
			log := GenerateRandomLog()
			fmt.Println(log)
		}
	}
}

// Generates a random log entry with a random log level.
func GenerateRandomLog() string {
	levels := []LogLevel{Info, Warning, Error, Debug}
	randomLevel := levels[rand.Intn(len(levels))]
	message := generateRandomMessage()
	// LOG FORMAT: [DATETIME] [LOG_LEVEL] [PID] message
	return fmt.Sprintf("[%s] [%s] [%v] %s", time.Now().Format(time.RFC3339), randomLevel,
		os.Getpid(),
		message)
}

// Generates a random log message
func generateRandomMessage() string {
	messages := []string{
		"This is an informational message.",
		"Something unusual happened.",
		"An error occurred while processing the request.",
		"Debugging information.",
	}
	return messages[rand.Intn(len(messages))]
}
