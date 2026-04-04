package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "username for login")
	var password *string = flag.String("password", "password", "password for login")
	var host *string = flag.String("host", "localhost", "hostname for database connection")
	var port *int = flag.Int("port", 3306, "port for database connection")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)

}
