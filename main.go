package main

import (
	"encoding/json"
	"github.com/oxakromax/Diffie-Hellman/DiffieHellman"
)

func main() {
	sucess := 0
	for i := 0; i < 10000; i++ {
		client, server := new(DiffieHellman.Diffie), new(DiffieHellman.Diffie)
		server.FirstConfig()             // Server First Configuration
		Sjson, _ := json.Marshal(server) // Server sends his config to Client
		err := client.JsonConfig(Sjson)  // Client reads server configs and set the commonKey
		if err != nil {
			return
		}
		//server.setComunnicatorKey(client.PublicKey)
		Cjson, _ := json.Marshal(client) // Client sends his config to the Server
		err = server.JsonConfig(Cjson)   // Server Reads the config of the Client and config the common Key
		if err != nil {
			return
		}
		if client.GetKey() == server.GetKey() {
			sucess++ // If the commonKey is Equal, sucess +1, to prove the 100% Sucess
		}
	}
	println(sucess)
}
