# Go F1 Library

A library to handle game packets from the F1 20xx game.

[![Pipeline](https://github.com/DaanV2/go-f1-library/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/DaanV2/go-f1-library/actions/workflows/pipeline.yaml)

# UPD example

```go
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
```