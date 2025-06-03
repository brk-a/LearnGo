package sniff_and_capture

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/user"
	"strconv"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // DO NOT use this in prod
}
var clients []*websocket.Conn
var mu sync.Mutex
var (
	iface    = ""
	snaplen  = int32(1600)
	promisc  = true
	timeout  = pcap.BlockForever
	devFound = false
)

func SniffAndCapture() {
	// 1. sniff, sniff
	if !is_root() {
		fmt.Println("user is not root")
		return
	}
	http.HandleFunc("/ws", ws_handler)
	go func() {
		fmt.Println("websocket server starting at port 4444")
		if err := http.ListenAndServe("0.0.0.0:4444", nil); err != nil {
			fmt.Println("error starting server", err)
		}
	}()

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(" ========== ID ==========  Name ========== ")
	for id, device := range devices {
		fmt.Printf(" ========== %d ==========  %s ========== \n", id, device)
	}

	var input string
	fmt.Scanln(&input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("invalid ID selected")
		return
	}
	for id, device := range devices {
		if id == number {
			devFound = true
			iface = device.Name
		}
	}

	if !devFound {
		log.Panicln("invalid interfac")
	}

	fmt.Printf("interface %s selected\n", iface)

	// 2. capture
	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}

	defer handle.Close()
	if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		log.Panicln(err)
	}
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("network traffic capture started ...")
	for packet := range source.Packets() {
		packet_str := packet.Dump()
		app_layer := packet.ApplicationLayer()
		if app_layer == nil {
			continue
		}
		payload := app_layer.Payload()
		if bytes.Contains(payload, []byte("uname")) || bytes.Contains(payload, []byte("passwd")) || bytes.Contains(payload, []byte("USER")) || bytes.Contains(payload, []byte("pass")) {
			fmt.Println(packet_str)
		}
		fmt.Println(" ========== ========== ========== ")
		broadcast_message(packet_str)
	}
}

func is_root() bool {
	current_user, err := user.Current()
	if err != nil {
		log.Fatalf("[is_root] unable to get current user: %s", err)
	}
	return current_user.Username == "root"
}

func ws_handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error upgrading", err)
		return
	}

	defer conn.Close()
	mu.Lock()
	clients = append(clients, conn)
	mu.Unlock()
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("client disconnected", err)
			break
		}
	}
	mu.Lock()
	for i, client := range clients {
		if client == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
	mu.Unlock()
}

func broadcast_message(message string) {
	mu.Lock()
	defer mu.Unlock()

	for _, conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			fmt.Println("error sending message", message)
		}
	}
}
