package main

import (
	"encoding/hex"
	"fmt"
	"lets-go-chat/pkg/hasher"
)

func main() {
	var hashed, _ = hasher.HashPassword("ololo")
	fmt.Println(hashed)

	var hash, _ = hex.DecodeString(hashed)
	var checked = hasher.CheckPasswordHash("ololo", hash)
	fmt.Println(checked)
}
