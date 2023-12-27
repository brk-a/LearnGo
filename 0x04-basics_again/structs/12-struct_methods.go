package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string{
	return fmt.Sprintf("Authorisation: Basic %v:%v", a.username, a.password)
}

func testAuth(authInfo authenticationInfo)  {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("==========================================================")
}

func main() {
	testAuth(authenticationInfo{
		username: "goatmatata",
		password: "12345",
	})
	testAuth(authenticationInfo{
		username: "bibuibui",
		password: "98765",
	})
	testAuth(authenticationInfo{
		username: "kakambwamwitu",
		password: "76921",
	})
}