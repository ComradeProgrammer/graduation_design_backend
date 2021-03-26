package git

import (
	"fmt"
	"os"
	"github.com/google/uuid"
)

type Git struct {
	SSHURL string
	Cwd    string
	FolderName string
	Users  []string
}

func NewGit(sshurl string, cwd string, clone bool) (*Git, error) {
	var res Git = Git{}
	res.SSHURL = sshurl
	res.Cwd = cwd
	if clone {
		if sshurl == "" {
			return nil, fmt.Errorf("NewGit:sshurl not set")
		}
		os.Mkdir(cwd, 0777)
		res.FolderName=uuid.New().String()

		_, err := RunCommand("", "git", "clone", sshurl,res.FolderName)
		if err != nil {
			return nil, err
		}
	}
	return &res, nil
}

func (g *Git)getAllUser()error{
	
	return nil
}
