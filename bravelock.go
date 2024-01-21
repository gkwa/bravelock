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
	hardcodedStrategy := &filename.HardcodedStrategy{Filename: "example.txt"}
	reflectionStrategy := &filename.ReflectionStrategy{}

	fmt.Println("Hardcoded Strategy:", getFileNamingStrategy(hardcodedStrategy))
	fmt.Println("Reflection Strategy:", getFileNamingStrategy(reflectionStrategy))

	fileName := getFileNamingStrategy(reflectionStrategy)
	fileNameWithoutExt := filename.GetFnameWithoutExtension(fileName)

	fmt.Println("File name without extension:", fileNameWithoutExt)

	return nil
}

func getFileNamingStrategy(strategy filename.FilenameStrategy) string {
	return strategy.GetFilename()
}
