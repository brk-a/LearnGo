package network_credentials_sniffer

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var DevName = "lo"
var Found = false

func NetworkCredentialsSniffer() {
	// 1. find all network interfaces for packet capturing using pcap.FindAllDevs()
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln("unable to fetch network interfaces")
	}

	for _, if_dev := range devices {
		if if_dev.Name == DevName {
			Found = true
		}
	}

	// 2. if desired device is found, continue, otherwise, exit
	if !Found {
		log.Panicln("desired device not found")
	}

	// 3. open live capture handle on that network interface using pcap.OpenLive()
	handle, err := pcap.OpenLive(DevName, 1600, false, pcap.BlockForever)
	if err != nil {
		fmt.Print(err)
		log.Panicln("unable to open handle on device")
	}
	defer handle.Close()

	// 4. apply BPF (berkely packet filter) on new handle
	if err := handle.SetBPFFilter("tcp and port 21"); err != nil {
		log.Panicln(err)
	}

	// 5. display all filtered packets received on channel returned from gopacket.NewPacketSource()
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}
		data := appLayer.Payload()
		if bytes.Contains(data, []byte("USER")) || bytes.Contains(data, []byte("PASS")) {
			fmt.Println(string(data))
		}
	}
}
