package main

import "fmt"

const (
	PI_APPLICATION         = "ssh"
	PI_DEFAULT_CONFIG_PATH = "/etc/privacyidea/authorizedkeyscommand"
)

var (
	PI_DEFAULT_HEADERS = []string{"Content-Type: application/json", "Accept: application/json"}
)

func main() {
	fmt.Println("piakc v0.1.0")
}
