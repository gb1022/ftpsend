package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"strconv"
	"time"

	"net"

	"github.com/goftp"
	"golang.org/x/crypto/ssh"
)

type Conf struct {
	SendPath1 string
	SendPath2 string
	SendPath3 string
	Ip        string
	F_port    string
	S_port    string
	User      string
	Password  string
	Cmd1      string
	Cmd2      string
}

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以 ,不用的话会报错
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

func byteToString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func ReadConf() Conf {
	var conf_data Conf
	data, err := ioutil.ReadFile("./conf.txt")
	if err != nil {
		return conf_data
	}
	//	fmt.Println("data:\n", byteToString(data))
	for _, line := range strings.Split(byteToString(data), "\n") {
		l := strings.Split(line, "=")
		num := len(l[1])
		if l[0] == "sendpath1" {
			sendpath1 := l[1][0 : num-1]
			fmt.Println("sendpath1:", sendpath1)
			conf_data.SendPath1 = sendpath1
		} else if l[0] == "sendpath2" {
			sendpath2 := l[1][0 : num-1]
			fmt.Println("sendpath2:", sendpath2)
			conf_data.SendPath2 = sendpath2
		} else if l[0] == "sendpath3" {
			sendpath3 := l[1][0 : num-1]
			fmt.Println("sendpath3:", sendpath3)
			conf_data.SendPath3 = sendpath3
		} else if l[0] == "ip" {
			ip := l[1][0 : num-1]
			fmt.Println("ip:", ip)
			conf_data.Ip = ip
		} else if l[0] == "f_port" {
			f_port := l[1][0 : num-1]
			fmt.Println("f_port:", f_port)
			conf_data.F_port = f_port
		} else if l[0] == "s_port" {
			s_port := l[1][0 : num-1]
			fmt.Println("s_port:", s_port)
			conf_data.S_port = s_port
		} else if l[0] == "user" {
			user := l[1][0 : num-1]
			fmt.Println("user:", user)
			conf_data.User = user
		} else if l[0] == "password" {
			password := l[1][0 : num-1]
			fmt.Println("password:", password)
			conf_data.Password = password
		} else if l[0] == "cmd1" {
			cmd1 := l[1][0 : num-1]
			fmt.Println("cmd1:", cmd1)
			conf_data.Cmd1 = cmd1
		} else if l[0] == "cmd2" {
			cmd2 := l[1]
			fmt.Println("cmd2:", cmd2)
			conf_data.Cmd2 = cmd2
		} else {
			fmt.Println("conffile is error!")
			return conf_data
		}
	}
	return conf_data

}

func GetPath() string { //获取当前路径
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	fmt.Println("getpath :", path)
	return path
}

func GetFiles(path string) []string { //获取文件
	//	path = "F:\\策划\\越南\\开发版本（M10-2）\\json_vietn"
	filelist := []string{}
	fmt.Println("path:", path)
	dir_list, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	for _, v := range dir_list {
		if v.IsDir() {
			continue
		}
		filelist = append(filelist, v.Name())
		//		fmt.Println("file:", v.Name())
	}

	return filelist
}

func SendFile(conf_data Conf, ftp *goftp.FTP, files []string, path_ex string, sport int) error {
	var err error
	var session *ssh.Session
	session, err = connect(conf_data.User, conf_data.Password, conf_data.Ip, sport)
	if err != nil {
		return err
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	cmd := ""
	for _, f := range files {
		sendFilePath1 := conf_data.SendPath1 + f
		delcmd := "rm -rf " + sendFilePath1 + ";"
		cmd += delcmd
	}
	fmt.Println("cmd:", cmd)
	session.Run(cmd)
	for _, f := range files {

		var file *os.File
		filepath := path_ex + f
		if file, err = os.Open(filepath); err != nil {
			return err
		}
		if conf_data.SendPath1 != "" {
			sendFilePath1 := conf_data.SendPath1 + f
			if err := ftp.Stor(sendFilePath1, file); err != nil {
				return err
			}
			fmt.Println("send ok file1:", sendFilePath1)
		}
		//		if conf_data.SendPath2 != "" {
		//			sendFilePath2 := conf_data.SendPath2 + f
		//			if err := ftp.Stor(sendFilePath2, file); err != nil {
		//				return err
		//			}
		//			fmt.Println("send ok file2:", sendFilePath2)
		//		}
		//		if conf_data.SendPath3 != "" {
		//			sendFilePath3 := conf_data.SendPath3 + f
		//			if err := ftp.Stor(sendFilePath3, file); err != nil {
		//				return err
		//			}
		//			fmt.Println("send ok file3:", sendFilePath3)
		//		}
	}
	for _, f := range files {

		var file *os.File
		filepath := path_ex + f
		if file, err = os.Open(filepath); err != nil {
			return err
		}
		//		if conf_data.SendPath1 != "" {
		//			sendFilePath1 := conf_data.SendPath1 + f
		//			if err := ftp.Stor(sendFilePath1, file); err != nil {
		//				return err
		//			}
		//			fmt.Println("send ok file1:", sendFilePath1)
		//		}
		if conf_data.SendPath2 != "" {
			sendFilePath2 := conf_data.SendPath2 + f
			if err := ftp.Stor(sendFilePath2, file); err != nil {
				return err
			}
			fmt.Println("send ok file2:", sendFilePath2)
		}
		//		if conf_data.SendPath3 != "" {
		//			sendFilePath3 := conf_data.SendPath3 + f
		//			if err := ftp.Stor(sendFilePath3, file); err != nil {
		//				return err
		//			}
		//			fmt.Println("send ok file3:", sendFilePath3)
		//		}
	}
	for _, f := range files {

		var file *os.File
		filepath := path_ex + f
		if file, err = os.Open(filepath); err != nil {
			return err
		}
		//		if conf_data.SendPath1 != "" {
		//			sendFilePath1 := conf_data.SendPath1 + f
		//			if err := ftp.Stor(sendFilePath1, file); err != nil {
		//				return err
		//			}
		//			fmt.Println("send ok file1:", sendFilePath1)
		//		}
		//		if conf_data.SendPath2 != "" {
		//			sendFilePath2 := conf_data.SendPath2 + f
		//			if err := ftp.Stor(sendFilePath2, file); err != nil {
		//				return err
		//			}
		//			fmt.Println("send ok file2:", sendFilePath2)
		//		}
		if conf_data.SendPath3 != "" {
			sendFilePath3 := conf_data.SendPath3 + f
			if err := ftp.Stor(sendFilePath3, file); err != nil {
				return err
			}
			fmt.Println("send ok file3:", sendFilePath3)
		}
	}
	return nil
}

func main() {
	var err error
	var ftp *goftp.FTP
	var session *ssh.Session
	conf_data := ReadConf()
	path := GetPath()
	path_ex := path + "..\\"
	files := GetFiles(path_ex)
	addr := conf_data.Ip + ":" + conf_data.F_port
	fmt.Println("SSH begin...")
	sport, _ := strconv.Atoi(conf_data.S_port)
	session, err = connect(conf_data.User, conf_data.Password, conf_data.Ip, sport)
	if err != nil {
		panic(err)
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	fmt.Println("remote print:", conf_data.Cmd1)
	session.Run(conf_data.Cmd1)
	fmt.Println("cmd1 do end!")
	if ftp, err = goftp.Connect(addr); err != nil {
		panic(err)
	}

	defer func() {
		ftp.Close()
		session.Close()
	}()
	fmt.Println("FTP Successfully connected !!")
	if err = ftp.Login(conf_data.User, conf_data.Password); err != nil {
		panic(err)
	}
	fmt.Println("FTP Successfully login !!")
	if err = SendFile(conf_data, ftp, files, path_ex, sport); err != nil {
		panic(err)
	}
	fmt.Println("FTP Successfully All post")
	fmt.Println("--------------------------------------")
	session, err = connect(conf_data.User, conf_data.Password, conf_data.Ip, sport)
	if err != nil {
		panic(err)
	}
	fmt.Println("remote print:", conf_data.Cmd2)
	session.Run(conf_data.Cmd2)
	fmt.Println("cmd2 do end!")
	time.Sleep(5 * time.Second)
}
