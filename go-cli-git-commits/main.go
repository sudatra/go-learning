package main

import (
	"flag"
)

func scan(path string) {
	print("scan");
}

func stats(email string) {
	print("stats");
}

func main() {
	var folder string;
	var email string;

	flag.StringVar(&folder, "add", "", "Add a folder to scan for git repos");
	flag.StringVar(&email, "email", "your@email.com", "The email ID to scan for");
	flag.Parse();

	if folder != "" {
		scan(folder);
		return;
	}

	stats(email);
}