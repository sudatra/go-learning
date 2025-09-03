package scanner

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/user"
)

func getDotFilePath() string {
	usr, err := user.Current();
	if err != nil {
		log.Fatal(err);
	}

	dotFile := usr.HomeDir + "/.goGitLocalStats";
	return dotFile
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
		existingRepos := parseFileLinesToSlice(filePath);
		repos := joinSlices(newRepos, existingRepos);

		dumpStringsSliceToFile(repos, filePath);
}

func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath);
	defer f.Close();
	
	var lines []string;
	scanner := bufio.NewScanner(f);

	for scanner.Scan() {
		lines = append(lines, scanner.Text());
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err);
		}
	}

	return lines;
}


func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)

			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}