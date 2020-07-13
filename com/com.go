package com

import(
	"os"
	"os/exec"
)

func command(c string) {
	cmd := exec.Command("cmd","/c",c)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Build() {
	command("go build")
}
func Cls() {
	command("cls")
}
func Code() {
	command("code .")
}
func Dir() {
	command("dir")
}
