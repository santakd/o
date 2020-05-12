package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestVimInfo(t *testing.T) {
	LoadVimLocationHistory(expandUser(vimLocationHistoryFilename))
}

func TestEmacsPlaces(t *testing.T) {
	LoadEmacsLocationHistory(expandUser(emacsLocationHistoryFilename))
}

func TestNeoVimMsgPack(t *testing.T) {
	curdir, err := os.Getwd()
	if err != nil {
		t.Fail()
	}
	searchFilename, err := filepath.Abs(filepath.Join(curdir, "main.go"))
	if err != nil {
		t.Fail()
	}
	line, err := FindInNvimLocationHistory(expandUser(nvimLocationHistoryFilename), searchFilename)
	if err != nil {
		// main.go might not be in the neovim location history, this is fine
		fmt.Println(err)
	}
	fmt.Println("line", line)
}
