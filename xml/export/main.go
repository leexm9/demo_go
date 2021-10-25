package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIp   string `xml:"serverIp"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName: "Shanghai_VPN", ServerIp: "127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName: "Beijing_VPN", ServerIp: "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
