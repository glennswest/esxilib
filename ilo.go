package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	//"os"
        "log"
	//"strconv"
	"strings"
	"time"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {

var servname string

        servname = "none"
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
                log.Printf("CMD = %s",temp)
                co := strings.Fields(temp)
                switch(co[0]){
                   case "USER_LOGIN":
                        log.Printf("Logging into: %s",co[1])
                        servname = co[1]
                        break
                   case "xml":
                        log.Printf("xml version requested")
                        result := "<RIBCL VERSION=\"2.0\"></RIBCL>"
                        send_result(c,result)
                        break
                   case "GET_FW_VERSION":
                        log.Printf("FW Version requested")
                        result := "<GET_FW_VERSION\r\n FIRMWARE_VERSION=\"1.91\"\r\n MANAGEMENT_PROCESSOR=\"2.22\"\r\n />"
                        send_result(c,result)
                        break
                   case "GET_HOST_POWER_STATUS":
                        log.Printf("Host Power Status for %s",servname)
                        break
                   case "SET_HOST_POWER":
                        log.Printf("Set Host Power Status for %s: %s",servname,co[1])
                        break
                   default: 
                        log.Printf("Unsupported command: %s", co[0])
                        break
                   }
	}
	c.Close()
}

func send_result(c net.Conn, thevalue string){
     c.Write([]byte(string(thevalue)))
}

func main() {
	PORT := ":1234"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
                log.Printf("New Connection")
		go handleConnection(c)
	}
}


