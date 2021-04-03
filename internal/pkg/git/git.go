package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Git struct {
	SSHURL     string
	FolderName string
}
type GitBlameLineCallback func(filename string, blame string)
type GitCommitCallback func(gitlog string)

func NewGit(sshurl string, cwd string, clone bool) (*Git, error) {
	var res Git = Git{}
	res.SSHURL = sshurl
	if clone {
		if sshurl == "" {
			return nil, fmt.Errorf("NewGit:sshurl not set")
		}
		if cwd != "" {
			os.Mkdir(cwd, 0777)
		}

		if cwd == "" {
			res.FolderName = uuid.New().String()
		} else {
			res.FolderName = cwd + "/" + uuid.New().String()
		}

		_, err := RunCommand("", "git", "clone", sshurl, res.FolderName)
		if err != nil {
			return nil, err
		}
	}
	return &res, nil
}

func (g *Git) GetAllTrackedFiles() ([]string, error) {
	//git ls-tree -r master --name-only
	res, err := RunCommand(g.FolderName, "git", "ls-tree", "-r", "master", "--name-only")
	if err != nil {
		return nil, err
	}
	return strings.Split(res, "\n"), nil
}

func (g *Git) BlameAllFile(files []string, callbacks ...GitBlameLineCallback) error {
	for _, file := range files {
		res, err := RunCommand(g.FolderName, "git", "blame", "-f", file)
		if err != nil {
			return err
		}
		split := strings.Split(res, "\n")
		for _, line := range split {
			for _, callback := range callbacks {
				callback(file, line)
			}

		}
	}
	return nil
}

func (g *Git) ReadAllCommit(callbacks ...GitCommitCallback) error {
	// command is  git log --pretty=format:"%H %an %ai %s" --shortstat
	//don't know why " needs to be removed but it works
	res, err := RunCommand(g.FolderName, "git", "log", "--pretty=format:%H %an %ai %s", "--shortstat")
	if err != nil {
		return err
	}
	strs:=strings.Split(res,"\n")
	for i:=0;i<len(strs);i=i+3{
		commitstr:=strs[i]+"\n"
		if i+1<len(strs){
			commitstr+=strs[i+1]+"\n"
		}
		if i+2<len(strs){
			commitstr+=strs[i+2]
		}
		for _,callback:=range callbacks{
			callback(commitstr)
		}
	}
	return nil
}

func (g *Git) Clear() error {
	if g.FolderName != "" {
		_, err := RunCommand("", "rm", "-rf", g.FolderName)
		return err
	}
	return nil
}
