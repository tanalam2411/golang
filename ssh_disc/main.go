package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"os"
)


type SSHCommand struct {
	Path string
	Env []string
	Stdin io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
	var (
		session *ssh.Session
	)
}

func main() {

	username := "root"
	password := "FixStream"

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	client := &SSHClient{
		Config: sshConfig,
		Host:   "172.16.2.175",
		Port:   22,
	}

	cmd := &SSHCommand{
		Path:   "hostname",
		Env:    []string{"LC_DIR=/"},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
