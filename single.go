package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func pullRepo(c *cli.Context) error {
	for _, arg := range c.Args().Slice() {
		_ = goclone(arg)
	}
	return nil
}

func goclone(gitSshURL string) error {
	// cloneUrl:
	// https://github.com/keaising/goclone
	// git@github.com:keaising/goclone.git
	// https://github.com/keaising/goclone.git
	urlFragments := splitURLIntoFragments(gitSshURL)

	direcotry, err := getTargetDir()
	if err != nil {
		return nil
	}
	targetDir := path.Join(direcotry, "src", urlFragments["host_name"], urlFragments["org"], urlFragments["repo"])
	fmt.Println("targetDir:", targetDir)
	gitSshURL = fmt.Sprintf("https://%s/%s/%s.git", urlFragments["host_name"], urlFragments["org"], urlFragments["repo"])

	return cloneRepo(gitSshURL, targetDir)
}

func getTargetDirectory() (string, error) {
	directory := fetch("git config clone.directory")
	directory = strings.Trim(directory, "\r\n")
	directory = strings.TrimSpace(directory)
	if directory == "" {
		return "", fmt.Errorf("Please set clone directory by 'git config --global clone.directory YOUR.PATH'")
	}
	homes := []string{"$HOME", "~", "HOME"}
	for _, home := range homes {
		if strings.HasPrefix(directory, home) {
			userHome, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			directory = path.Join(userHome, strings.TrimPrefix(directory, home))
			return directory, nil
		}
	}
	return directory, nil
}

func getTargetDir() (string, error) {
	path := os.Getenv("GOPATH")
	fmt.Println(path)
	return path, nil
}

type urlFragments struct {
	host, org, repo string
}

func splitURLIntoFragments(url string) (urlFragments map[string]string) {
	regStr := `[a-zA-Z]+(@|://)(?P<host_name>\w[\w.-]+)(:|/)(?P<org>\w[\w.-]+)/(?P<repo>\w[\w.-]+)`
	if strings.HasPrefix(url, "git") {
		regStr = `[a-zA-Z]+(@|://)(?P<host_name>\w[\w.-]+)(:|/)(?P<org>\w[\w.-]+)/(?P<repo>\w[\w.-]+).git`
	}
	compRegEx := regexp.MustCompile(regStr)
	match := compRegEx.FindStringSubmatch(url)

	urlFragments = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			urlFragments[name] = match[i]
		}
	}
	return
}

func cloneRepo(cloneUrl string, directory string) error {
	// git clone git@github.com:keaising/goclone.git /Users/mac/source/test/keaising/goclone
	cmd := fmt.Sprint("git clone ", cloneUrl, " ", directory)
	fmt.Println(cmd)

	_, err := os.Stat(directory)
	if !os.IsNotExist(err) {
		err = os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return execute(cmd)
}
