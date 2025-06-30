package main

import (
	advanced_port_scanner "github.com/brk-a/offensive_go/advanced_port_scanner"
	net_cred_sniffer "github.com/brk-a/offensive_go/network_credentials_sniffer"
	port_scanner "github.com/brk-a/offensive_go/port_scanner"
	remote_admin_tool "github.com/brk-a/offensive_go/remote_administration_tool"
	remote_shell "github.com/brk-a/offensive_go/remote_shell"
	sniff_n_capture "github.com/brk-a/offensive_go/sniff_and_capture"
	sub_domain_finder "github.com/brk-a/offensive_go/sub_domain_finder"
	key_logger "github.com/brk-a/offensive_go/web_key_logger"
	titan_stealer "github.com/brk-a/offensive_go/titan_stealer"
)

func main() {
	port_scanner.Main()
	remote_shell.Main()
	advanced_port_scanner.Main()
	sniff_n_capture.Main()
	key_logger.Main()
	sub_domain_finder.Main()
	net_cred_sniffer.Main()
	remote_admin_tool.Main()
	titan_stealer.Main()
}
