package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	f1_2023 "github.com/DaanV2/go-f1-library/packets/2023"
	test_util "github.com/DaanV2/go-f1-library/test/packets/util"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", ":21200")
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on", udpAddr)

	results := test_util.NewResults()
	parser := f1_2023.NewPacketParser()
	defer func() {
		err := results.Store(path.Join(".", "test", "packets", "2023"))
		if err != nil {
			fmt.Println(err)
		}
	}()


	// Read from UDP listener in endless loop
	go func() {
		var buf [f1_2023.MAX_PACKET_SIZE]byte
		for {
			n, _, err := conn.ReadFromUDP(buf[0:])
			if err != nil {
				// Conn closed
				if strings.Contains(err.Error(), "use of closed network connection"){
					return
				}

				fmt.Println(err)
			} else {
				results.AddPacket(parser, buf[:n])
			}
		}
	}()

	// Listen for ctrl-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	conn.Close()
}


