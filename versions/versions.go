package versions

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

type Version struct {
	Major         int
	Minor         int
	Build         int
	Repo          string
	Branch        string
	ShortCommitId string
	LongCommitId  string
}

type Generator interface {
	Generate(v Version, filename string)
}

var verbose = false

func GetBranchWithHead(p string) string {

	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	temppath := path.Join(wd, p)
	if verbose {
		fmt.Printf("Changing dir to %s\n", temppath)
	}
	os.Chdir(temppath)
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}").Output()
	if err != nil {
		if verbose {
			fmt.Printf("Unable to determine branch for %s\n", p)
			return "unknown"
		}
	}
	if verbose {
		fmt.Printf("repo version %s", out)
	}
	outstr := strings.TrimSpace(string(out))
	sp := strings.Split(outstr, "/")
	return sp[1]
}

// Function GetCommitId
//          returns the pair of short commid id and the full long commit id
//
//          p = path to change to
//          rev = if the commit id cannot be found, return this as a pair
func GetCommitId(p string, rev string) (string, string) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	temppath := path.Join(wd, p)
	if verbose {
		fmt.Printf("Changing working dir to %s\n", temppath)
	}
	os.Chdir(temppath)
	if verbose {

	}
	//out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	out, err := exec.Command("git", "log", "-n", "1", "--pretty=format:%h").Output()
	if err != nil {
		if verbose {
			fmt.Printf("Error %v getting commit id\n", err)
		}
		return rev, rev
	}
	shortid := strings.TrimSpace(string(out))
	out, err = exec.Command("git", "log", "-n", "1", "--pretty=format:%H").Output()
	if err != nil {
		if verbose {
			fmt.Printf("Error %v getting commit id\n", err)
		}
		return rev, rev
	}
	fullid := strings.TrimSpace(string(out))

	return shortid, fullid
}

func GetRemoteURL(p string) string {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	temppath := path.Join(wd, p)
	if verbose {
		fmt.Printf("Changing working dir to %s\n", temppath)
	}
	os.Chdir(temppath)
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		if verbose {
			fmt.Printf("Error %v getting commit id\n", err)
		}
		return "?"
	}
	remote := string(out)
	sepexp := regexp.MustCompile("[^\\s]+")
	remfields := sepexp.FindAllString(remote, -1)
	remrepo := "?"
	if remfields[5] == "(push)" {
		remrepo = remfields[4]
	}
	return remrepo
}

func Report() {
	fmt.Printf("Version : %d.%d-%d\n", versionMajor, versionMinor, versionBuild)
	fmt.Printf("Built : %s\n", buildTime)
	fmt.Printf("Repo URL : %s\n", repoURL)
	fmt.Printf("Branch : %s\n", branchName)
	fmt.Printf("Commit Id : Short : %s Long %s\n", shortCommitId, longCommitId)
}
