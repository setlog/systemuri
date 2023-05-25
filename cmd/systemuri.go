package main

import "github.com/setlog/systemuri"

func main() {
	systemuri.RegisterURLHandler("Test", "testtest", "testapplication")
}
