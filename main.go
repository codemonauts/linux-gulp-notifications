package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gen2brain/beeep"
)

type UDPData struct {
	Title       string `json:"title"`
	Information string `json:"information"`
	Error       bool   `json:"error"`
}

var iconLookup = map[bool]string{
	true:  "icons/error.png",
	false: "icons/info.png",
}

func main() {
	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 9090, Zone: ""})
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, _, _ := ServerConn.ReadFromUDP(buf)
		message := UDPData{}
		err := json.Unmarshal(buf[0:n], &message)
		if err != nil {
			panic(err)
		}
		fmt.Println(message.Error)

		err = beeep.Notify(message.Title, message.Information, iconLookup[message.Error])
		if err != nil {
			panic(err)
		}

	}
}
