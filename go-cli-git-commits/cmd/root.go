package cmd

import (
	"flag"
	"github.com/sudatra/go-cli-git-commits/scanner"
	"github.com/sudatra/go-cli-git-commits/stats"
)

func Execute() {
	var folder string;
	var email string;

	flag.StringVar(&folder, "add", "", "Add a folder to scan for git repos");
	flag.StringVar(&email, "email", "your@email.com", "The email ID to scan for");
	flag.Parse()

	if folder != "" {
		scanner.Scan(folder);
		return;
	}

	stats.Stats(email);
}