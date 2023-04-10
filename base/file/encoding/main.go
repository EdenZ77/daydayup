package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	//api := "dWnFh6rWw9GAVuEp1A77Pgf2CpTSwmsJ"
	api := "dWnFh"
	toString := base64.StdEncoding.EncodeToString([]byte(api))
	fmt.Println(toString)

	toString1 := base64.URLEncoding.EncodeToString([]byte(api))
	fmt.Println(toString1)

	toString2 := base64.RawStdEncoding.EncodeToString([]byte(api))
	fmt.Println(toString2)

	toString3 := base64.RawURLEncoding.EncodeToString([]byte(api))
	fmt.Println(toString3)
}
