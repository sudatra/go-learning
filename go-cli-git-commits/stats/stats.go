package stats

import (
	"fmt"
	"sort"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/sudatra/go-cli-git-commits/scanner"
)

const daysInLastSixMonths = 183;
const outOfRange = 99999;
type column []int;

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

func fillCommits(email string, path string, commits map[int]int) map[int]int {
	repo, err := git.PlainOpen(path);
	if err != nil {
		panic(err);
	}

	ref, err := repo.Head();
	if err != nil {
		panic(err);
	}

	iterator,err := repo.Log(&git.LogOptions{From: ref.Hash()});
	if err != nil {
		panic(err);
	}

	offset := calcOffset();
	err = iterator.ForEach(func(c *object.Commit) error {
		daysAgo := countDaysSinceDate(c.Author.When) + offset;

		if c.Author.Email != email {
			return nil;
		}

		if daysAgo != outOfRange {
			commits[daysAgo]++;
		}

		return nil;
	})

	if err != nil {
		panic(err);
	}

	return commits;
}

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date();
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location());
	
	return startOfDay;
}

func countDaysSinceDate(date time.Time) int {
	days := 0;
	now := getBeginningOfDay(time.Now());

	for date.Before(now) {
		date = date.Add(time.Hour * 24);
		days++;

		if days > daysInLastSixMonths {
			return outOfRange;
		}
	}

	return days;
}

func calcOffset() int {
	var offset int;
	weekday := time.Now().Weekday();

	switch weekday {
	case time.Sunday:
		offset = 7;
	case time.Monday:
		offset = 6;
	case time.Tuesday:
		offset = 5;
	case time.Wednesday:
		offset = 4;
	case time.Thursday:
		offset = 3;
	case time.Friday:
		offset = 2;
	case time.Saturday:
		offset = 1;
	}

	return offset;
}

func printCommitStats(commits map[int]int) {
	keys := sortMapIntoSlice(commits);
	cols := buildCols(keys, commits);

	printCells(cols);
}

func sortMapIntoSlice(mp map[int]int) []int {
	var keys []int;
	for k := range m {
		keys = append(keys, k);
	}

	sort.Ints(keys);
	return keys;
}

func buildCols(keys []int, commits map[int]int) map[int]column {
	cols := make(map[int]column);
	col := column{};

	for _, k := range keys {
		week := int(k / 7);
		dayInWeek := k % 7;

		if dayInWeek == 0 {
			col = column{};
		}

		col = append(col, commits[k]);
		if dayInWeek == 6 {
			cols[week] = col;
		}
	}

	return cols;
}

func printCells(cols map[int]column) {
	printMonths();

	for j := 6; j >= 0; j-- {
		for i := weeksInLastSixMonths + 1; i >= 0; i-- {
			if i == weeksInLastSixMonths + 1 {
				printDayCol(j);
			}

			if col, ok := cols[i]; ok {
				if i == 0 && j == calcOffset() - 1 {
					printCell(cols[j], true);
					continue;
				} else {
					if len(col) > j {
						printCell(col[j], false);
						continue;
					}
				}
			}

			printCell(0, false);
		}

		fmt.Printf("\n");
	}
}

func printDayCol(day int) {
	out := "		";
	switch day {
	case 1:
		out = " Mon ";
	case 3:
	out = " Wed ";
	case 5:
	out = " Fri ";
	}

	fmt.Printf(out);
}
