package main

// importing the required package for this code..

import (
	"fmt"
	"net"
	// "strconv"
	"context"

	"time"
	"github.com/reiver/go-telnet"

)


// defining the global variable 

var (
	ipAddress = "172.190.22.156"
	port      = 22
	// timeout   = 5 * time.Second
)


// main function to call the other functions
func main() {
	
	// calling the checkSSHConnection function.

	sshEnabled, result := checkSSHConnection(ipAddress, port)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	fmt.Println(result)
	fmt.Printf("SSH connection enabled: %t\n", sshEnabled)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
}
	


// define the function to check the ssh connection is enable to the perticular VM

func checkSSHConnection(ip string, port int) (bool, string) {
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Use the context to establish the Telnet connection
	dialer := &telnet.Dialer{}

	conn, err := dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", ip, port))

	if err != nil {

		return false, fmt.Sprintf("SSH connection to %s:%d failed: %s", ip, port, err.Error())
	}

	defer conn.Close()

	return true, fmt.Sprintf("SSH connection to %s:%d successful", ip, port)
}




