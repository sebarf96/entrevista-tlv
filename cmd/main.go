package main

import (
	fun "entrevista/fif/sebita/internal/functions"
	"fmt"
	"log"
)

func main() {
	var tlv string
	fmt.Println("WELCOME!-------------------------")
	fmt.Println("Inserta el tlv: ")
	fmt.Scanln(&tlv)
	res, err := fun.TlvToMAP([]byte(tlv))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MAP = %s \n", res)
}
