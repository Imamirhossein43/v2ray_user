package pkg

import "fmt"

type Server1 struct {
	ID      int64
	Address string
	Port    int
}

func Server() {
	
	server := Server1{
		ID:      1,
		Address: "http://87.248.150.177",
		Port:    443,
	}
	fmt.Println(server)
}
