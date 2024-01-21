package bravelock

import (
	"fmt"
	"log/slog"

	"github.com/jessevdk/go-flags"
	"github.com/taylormonacelli/bravelock/filename"
)

var opts struct {
	LogFormat string `long:"log-format" choice:"text" choice:"json" default:"text" description:"Log format"`
	Verbose   []bool `short:"v" long:"verbose" description:"Show verbose debug information, each -v bumps log level"`
	logLevel  slog.Level
}

func Execute() int {
	if err := parseFlags(); err != nil {
		slog.Error("error parsing flags", "error", err)
		return 1
	}

	if err := setLogLevel(); err != nil {
		slog.Error("error setting log level", "error", err)
		return 1
	}

	if err := setupLogger(); err != nil {
		slog.Error("error setting up logger", "error", err)
		return 1
	}

	if err := run(); err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}

	return 0
}

func parseFlags() error {
	_, err := flags.Parse(&opts)
	if err != nil {
		return fmt.Errorf("parse flags failed: %w", err)
	}
	return nil
}

func run() error {
	var fileName string

	hardcodedStrategy := &filename.HardcodedStrategy{Filename: "example.txt"}
	reflectionStrategy := &filename.ReflectionStrategy{}

	fileName = hardcodedStrategy.GetFilename()
	fmt.Println("Hardcoded Strategy:", fileName)

	fileName = reflectionStrategy.GetFilename()
	fmt.Println("Reflection Strategy:", fileName)

	fileNameWithoutExt := filename.GetFnameWithoutExtension(fileName)
	fmt.Println("File name without extension:", fileNameWithoutExt)

	return nil
}
