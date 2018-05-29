package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	repoPath := flag.String("repopath", "", "")

	flag.Parse()

	diff, err := getDiffFromRemote(*repoPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	prettyPrint(diff)
}

func prettyPrint(diff []string) {
	for _, e := range diff {
		fmt.Println(e)
	}
}

func getDiffFromRemote(repoPath string) ([]string, error) {
	orig, err := getSortedCommits(repoPath)
	if err != nil {
		return nil, err
	}

	if err := fetch(repoPath); err != nil {
		return nil, err
	}

	new, err := getSortedCommits(repoPath)
	if err != nil {
		return nil, err
	}

	return getDiff(orig, new), nil
}

func getGitDir(repoPath string) string {
	return fmt.Sprintf("%s/.git", repoPath)
}

func fetch(repoPath string) error {
	cmd := exec.Command("/usr/bin/git", "--git-dir", getGitDir(repoPath), "fetch", "--all")
	return cmd.Run()
}

func getHashMap(list []string) map[string]struct{} {
	result := make(map[string]struct{}, len(list))
	for _, e := range list {
		result[e] = struct{}{}
	}

	return result
}

func getDiff(orig, new []string) []string {
	origMap := getHashMap(orig)
	newMap := getHashMap(new)

	result := make([]string, 0)
	for e, _ := range newMap {
		if _, ok := origMap[e]; ok {
			continue
		}

		result = append(result, e)
	}

	return result
}

func getSortedCommits(repoPath string) ([]string, error) {
	cmd := exec.Command("/usr/bin/git", "--git-dir", getGitDir(repoPath), "rev-list", "--all")

	lines, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	commits := make([]string, 0, len(lines))
	for _, line := range bytes.Split(lines, []byte{'\n'}) {
		commits = append(commits, string(line))
	}

	return commits, nil
}
