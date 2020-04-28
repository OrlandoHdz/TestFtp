package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

	err = conftp.ChangeDir("/prod/gxs/Inbound")
	if err != nil {
		msg := fmt.Sprintf("Error al cambiar el directorio: %s", err)
		fmt.Printf(msg)
	}

	File, err := ioutil.ReadFile("866_1588060837-31285_4.EDI")
	if err != nil {
		fmt.Printf("Error al abrir archivo: %s\n", err)
	}

	mbytes := bytes.NewBuffer(File)

	err = conftp.Stor("prueba.edi", mbytes)
	if err != nil {
		fmt.Printf("Error al subir el archivo: %s\n", err)
	}
	fmt.Println("Archivo subido correctamente:")

	conftp.Quit()
}
