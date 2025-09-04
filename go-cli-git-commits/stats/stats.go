package stats

import (
	"fmt"
	"github.com/sudatra/go-cli-git-commits/scanner"
)

const daysInLastSixMonths = 183;

func processRepositories(email string) map[int]int {
	filePath := scanner.GetDotFilePath();
	repos := scanner.ParseFileLinesToSlice(filePath);
	daysInMap := daysInLastSixMonths;

	commits := make(map[int]int, daysInMap);
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0;
	}

	for _, path := range repos {
		commits = fillCommits(email, path, commits);
	}

	return commits;
}

func Stats(email string) {
	commits := processRepositories(email);
	printCommitStats(commits);
}