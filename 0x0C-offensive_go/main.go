package main

import (
	port_scanner "github.com/brk-a/offensive_go/port_scanner"
	remote_shell "github.com/brk-a/offensive_go/remote_shell"
	advanced_port_scanner "github.com/brk-a/offensive_go/advanced_port_scanner"
	sniff_n_capture "github.com/brk-a/offensive_go/sniff_and_capture"
	key_logger "github.com/brk-a/offensive_go/web_key_logger"
	sub_domain_finder "github.com/brk-a/offensive_go/sub_domain_finder"
	net_cred_sniffer "github.com/brk-a/offensive_go/network_credentials_sniffer"
)

func main(){
	port_scanner.Main()
	remote_shell.Main()
	advanced_port_scanner.Main()
	sniff_n_capture.Main()
	key_logger.Main()
	sub_domain_finder.Main()
	net_cred_sniffer.Main()
}