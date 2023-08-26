package main

import (
	"log"
	"os"
	"sync"

	"github.com/elliot40404/lv/pkg/db"
	"github.com/elliot40404/lv/pkg/server"
	"github.com/elliot40404/lv/pkg/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "lv",
		Usage:   "Log Visualizer - A simple CLI tool to visualize logs",
		Version: "0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Elliot",
				Email: "admin@elliot404.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Whether to run in debug mode",
			},
		},
		Action: func(*cli.Context) error {
			// Channel for the scanner to exit the server
			var exitChannel = make(chan bool)
			// Buffered channel for the server to send logs to the scanner
			var logChannel = make(chan string, 100)
			var wg sync.WaitGroup
			// if debug mode is enabled, set a ENV variable
			if cliFlags := os.Args[1:]; len(cliFlags) > 0 && cliFlags[0] == "--debug" {
				os.Setenv("DEBUG", "true")
			}
			// init in-memory db
			db.InitDB()
			wg.Add(1)
			go utils.ReadPipe(logChannel, exitChannel, &wg)
			wg.Add(1)
			go server.Run(logChannel, exitChannel, &wg)
			wg.Wait()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
