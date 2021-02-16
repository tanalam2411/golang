package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

/*
ssh server - https://github.com/gliderlabs/ssh
*/

func main() {

	//currentUser, err := user.Current()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Printf("current user: %v, home dir: %v\n",currentUser.Username, currentUser.HomeDir)

	config := &ssh.ClientConfig{
		User: "tan",
		Auth: []ssh.AuthMethod{
			ssh.Password("p"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "172.16.205.147:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	if err := session.Shell(); err != nil {
		panic(err)
	}

	if err := session.Wait(); err != nil {
		panic(err)
	}
}
