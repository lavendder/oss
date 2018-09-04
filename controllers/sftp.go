package controllers

import (
	"golang.org/x/crypto/ssh"
	"time"
	"fmt"
	"github.com/pkg/sftp"
	"net"
)
func connect(user, password, host string, port int) (*sftp.Client, error) {
	fmt.Println(222)
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	sshClient, err = ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		fmt.Println(err.Error(),1)

		return nil, err
	}

	// create sftp client
	sftpClient, err = sftp.NewClient(sshClient)
	if err != nil {
		fmt.Println(err.Error(),2)

		return nil, err
	}
	fmt.Println(222)
	return sftpClient, nil
}
