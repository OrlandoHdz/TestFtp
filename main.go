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

	conftp, err := ftp.Dial("0.0.0.0:21")
	if err != nil {
		log.Fatalf("Error ftp.Dial: %v", err)
	}
	err = conftp.Login("TEST", "*****")
	if err != nil {
		log.Fatalf("Error Login: %v", err)
	}

	err = conftp.ChangeDir("/test")
	if err != nil {
		msg := fmt.Sprintf("Error al cambiar el directorio: %s", err)
		fmt.Printf(msg)
	}

	File, err := ioutil.ReadFile("test.hk")
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
