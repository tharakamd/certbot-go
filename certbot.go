package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	log.Println("Hello this is certbot client manager")

	verifyCertbotClientExistance()

	cmd := exec.Command("C:\\Users\\Tharaka\\go\\src\\awesomeProject\\hello-world.exe", )
	cmd.Stdin = strings.NewReader("this is email this is name")
	var out bytes.Buffer
	cmd.Stdout = &out
	err2 := cmd.Run()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func verifyCertbotClientExistance(){
	_ , err := exec.LookPath("C:\\Users\\Tharaka\\go\\src\\awesomeProject\\hello-world.exe")
	if err != nil {
		log.Fatal("Can't find executable file for Certbot client")
	}
	log.Println("Certbot client located successfully")
}