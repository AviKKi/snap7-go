package main

import (
	"fmt"

	"github.com/avikki/snap7-go"
)

// func main() {
// 	snap7_client, err := snap7.ConnentTo("127.0.0.1", 0, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(snap7_client)
// }

func main() {
	host := "127.0.0.1"
	plc := snap7.NewLogo()
	err := plc.Connect(host, 0x0200, 0x0300, 8002)
	
	fmt.Println(plc)
	fmt.Println(err)
	plc.Read("V1.1")
	plc.Write("V1.1", 0)
}
