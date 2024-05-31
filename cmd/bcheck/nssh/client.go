package nssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func NClient(hostname string, port int, rootPwd string) *ssh.Client {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password(rootPwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", hostname+":"+strconv.Itoa(port), config)
	if err != nil {
		log.Fatal("unable to connect:", err)
	}
	return conn
}

func NoKeyClient(host string) (*ssh.Client, error) {
	user, _ := user.Current()
	currentHost, _ := os.Hostname()
	var (
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	key, err := ioutil.ReadFile(path.Join(homePath, ".ssh", "id_rsa"))
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	clientConfig = &ssh.ClientConfig{
		User: user.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if host == "" {
		addr = fmt.Sprintf("%s:%d", currentHost, 2343)

	}
	addr = fmt.Sprintf("%s:%d", host, 2343)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		err = errors.Wrapf(err, "")
		return nil, err
	}

	return client, nil
}
