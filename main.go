package main

import (
	"fmt"
	"log"

	"github.com/jlaffaye/ftp"
)

func main() {
	fmt.Println("Inicia proceso")

	conftp, err := ftp.Dial("172.16.100.40:21")
	if err != nil {
		log.Fatalf("Error ftp.Dial: %v", err)
	}
	err = conftp.Login("DSOHM", "SISAMEX.10")
	if err != nil {
		log.Fatalf("Error Login: %v", err)
	}

	conftp.Quit()
}
