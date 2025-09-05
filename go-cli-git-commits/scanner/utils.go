package scanner

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func GetDotFilePath() string {
	usr, err := user.Current();
	if err != nil {
		log.Fatal(err);
	}

	dotFile := usr.HomeDir + "/.goGitLocalStats";
	return dotFile
}

func ParseFileLinesToSlice(filePath string) []string {
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

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true;
		}
	}

	return false;
}

func joinSlices(newRepos []string, existingRepos []string) []string {
	for _, i := range newRepos {
		if !sliceContains(existingRepos, i) {
			existingRepos = append(existingRepos, i);
		}
	}

	return existingRepos;
}

func dumpStringsSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n");
	ioutil.WriteFile(filePath, []byte(content), 0755);
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
		existingRepos := ParseFileLinesToSlice(filePath);
		repos := joinSlices(newRepos, existingRepos);

		dumpStringsSliceToFile(repos, filePath);
}