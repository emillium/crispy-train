package chain

import (
	"fmt"
)

func (n PoSNetwork) PrintBlockchainInfo() {
	for i, block := range n.Blockchain {
		fmt.Println("Block ", i, " Info:")
		fmt.Println("\tTimestamp:", block.Timestamp)
		fmt.Println("\tPrevious Hash:", block.PrevHash)
		fmt.Println("\tHash:", block.Hash)
		fmt.Println("\tValidator Address:", block.ValidatorAddr)
	}
}
