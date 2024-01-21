package main

import (
	"os"

	"github.com/taylormonacelli/bravelock"
)

func main() {
	code := bravelock.Execute()
	os.Exit(code)
}
