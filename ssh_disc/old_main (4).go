package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	outDir   = `/home/tan/tanveer/golang/src/golang/ssh_disc/output`
	username = "root"
	password = "FixStream"
)

var count = 0
var totalIps = 0

var wg = sync.WaitGroup{}

func main() {

	startTime := time.Now()

	credentials := map[string]string{}
	credentials["a"] = "b"
	credentials["c"] = "d"
	credentials["e"] = "f"
	credentials["g"] = "h"
	credentials["root"] = "FixStream"

	_ = os.Mkdir(outDir, os.ModePerm)
	//commands := []string{"hostname", `ifconfig eth0 | head -2 | tail -1 | tr -s " " " " | cut -d ":" -f2 | cut -d " " -f1`,
	//	`dmidecode -t  0 | grep "Vendor"`, `dmidecode -t  1 | grep "Manufacturer"`, "uname -s", "uname -r", `ifconfig eth0 | grep HWaddr | tr -s " " " " | cut -d " " -f5`}

	commands := []string{"hostname", `free -m`, `df -h`, "ifconfig"}
	//commands := []string{"hostname"}

	subnets := []string{"172.16.2.0/28"}
	var ipList []string

	for _, subnet := range subnets {
		subnetIpList, err := Hosts(subnet)
		if err != nil {
			panic(err)
		}

		ipList = append(ipList, subnetIpList...)
	}

	fmt.Println("ip list: ", ipList)
	totalIps = len(ipList)

	for _, ip := range ipList {

		wg.Add(1)
		go sshExecutor(ip, username, password, commands, credentials)
	}

	fmt.Printf("\n\n\n\n ---------- GOMAXPROCS:%v \n\n\n\n\n\n\n ", runtime.GOMAXPROCS(-1))

	fmt.Println("--------- wg waiting -----------")
	wg.Wait()
	fmt.Println("--------- wg Done -----------")

	elapsed := time.Since(startTime)
	fmt.Printf("\n\n\n Total time %s", elapsed)

}

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)

	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func haha(ip string) {
	count += 1
	fmt.Println("hahahahaha: ", ip, count, totalIps)
}

func sshExecutor(ip, username, password string, commands []string, credentials map[string]string) error {

	defer wg.Done()
	defer haha(ip)

	client := ssh.Client{}
	_, _ = username, password

	retry := len(credentials)
	for uname, paswd := range credentials {
		_client, err := authenticate(ip, uname, paswd)
		client = _client
		retry -= 1

		fmt.Println(ip, uname, paswd, err)

		if err == nil {
			break
		}

		if retry == 0 {
			return err
		}

	}

	//client, err := authenticate(ip, username, password)
	//if err != nil {
	//	//log.Printf("unable to connect %v: %v", ip, err)
	//	//wg.Done()
	//	return err
	//}

	defer client.Close()

	// Create a session
	session, err := client.NewSession()
	if err != nil {
		//log.Println("Failed to create session: ", err)
		//wg.Done()
		return err
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Println(err)
		//wg.Done()
		return err
	}

	outputFilePath := filepath.Join(outDir, fmt.Sprintf("%v_output_1.txt", ip))

	if _, err := os.Stat(outputFilePath); err == nil {
		outputFilePath = filepath.Join(outDir, fmt.Sprintf("%v_output_2.txt", ip))

	}

	outFile, err := os.Create(outputFilePath)

	if err != nil {
		//log.Printf("Failed to create output file: %v", err)
		//wg.Done()
		return err
	}

	session.Stdout = outFile
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		log.Println(err)
		//wg.Done()
		return err
	}

	//for _, cmd := range os.Args[1:] {
	for _, cmd := range commands {
		//fmt.Println("Running cmd: ", cmd)
		stdin.Write([]byte(cmd + "\n"))

	}

	//fmt.Printf(" output: %#v \n\n ", session.Stdout)

	stdin.Write([]byte("logout\n"))
	//stdin.Write([]byte("N\n"))

	session.Wait()

	//wg.Done()
	//fmt.Println("Done: ", ip)
	return nil
}

func authenticate(host, username, password string) (ssh.Client, error) {
	//hostKey, err := checkHostKey(host)
	//if err != nil {
	//	fmt.Println(err)
	//	//log.Fatal(err)
	//}

	//key, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"))
	//if err != nil {
	//	log.Fatalf("unable to read private key: %v", err)
	//}

	// Create the Signer for this private key.
	//signer, err := ssh.ParsePrivateKey(key)
	//if err != nil {
	//	log.Fatalf("unable to parse private key: %v", err)
	//}

	// Create client config
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			// Use the PublicKeys method for remote authentication.
			//ssh.PublicKeys(signer),
		},
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 1,
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", host+":22", config)

	if err != nil {
		return ssh.Client{}, err
	}

	return *client, err
}

func checkHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}

		if strings.Contains(fields[0], host) {

			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
	}

	return hostKey, nil
}
