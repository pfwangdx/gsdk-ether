package main

import "fmt"

func main() {

	// new client
	url := "http://localhost:8545"
	// url := "/home/wpf/go/src/github.com/DxChainNetwork/run-main-node/data/gdx.ipc"
	rawClient, err := Dial(url)
	if err != nil {
		panic(err)
	}

	// get block number
	var result string
	err = rawClient.GetBlockNumber(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("get block number result", result)

	rawClient.Close()
}
