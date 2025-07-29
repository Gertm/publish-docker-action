package main

import (
	"fmt"
	"log/slog"
	"os"
)

func setOutput(name, value string) {
	f, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("cannot open GITHUB_OUTPUT for writing", "error", err)
		return
	}
	_, err = fmt.Fprintf(f, "{%s}={%s}", name, value)
	if err != nil {
		slog.Error("cannot write to GITHUB_OUTPUT", "error", err)
	}
}
