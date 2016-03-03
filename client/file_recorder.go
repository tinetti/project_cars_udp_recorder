package client

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

/* A Simple function to verify error */
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	/* Prepare address at port 5606 */
	serverAddr, err := net.ResolveUDPAddr("udp", ":5606")
	checkError(err)

	/* Now listen at selected port */
	serverConn, err := net.ListenUDP("udp", serverAddr)
	checkError(err)
	defer serverConn.Close()
	fmt.Println("Started listening on port", serverAddr)

	buf := make([]byte, 2048)
	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Println("Received", n, "bytes from addr", addr)

		filename := "/tmp/project_cars/" + time.Now().String()
		mode := os.FileMode(0644)
		slice := buf[:n]
		err = ioutil.WriteFile(filename, slice, mode)
		checkError(err)

		fmt.Println("Wrote", filename)
	}
}
