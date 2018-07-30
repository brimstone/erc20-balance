package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/brimstone/go-erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getContract(address string, client *ethclient.Client) *ERC20.ERC20 {
	account := common.HexToAddress(address)
	instance, err := ERC20.NewERC20(account, client)
	if err != nil {
		panic(err)
	}
	return instance
}

func prettyPrint(balance *big.Int, power int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(power)))
}

func main() {

	clientPtr := flag.String("node", "http://localhost:8545", "Address of Ethereum node")
	accountPtr := flag.String("account", "", "Address of account with tokens")
	contractPtr := flag.String("contract", "", "Address of contract with tokens")

	flag.Parse()

	if *accountPtr == "" {
		log.Fatal("Account must be set!")
	}
	if *contractPtr == "" {
		log.Fatal("Contract must be set!")
	}
	client, err := ethclient.Dial(*clientPtr)
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress(*accountPtr)

	instance := getContract(*contractPtr, client)

	symbol, err := instance.Symbol(nil)
	if err != nil {
		panic(err)
	}

	decimals, err := instance.Decimals(nil)
	if err != nil {
		panic(err)
	}

	balance, err := instance.BalanceOf(nil, account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %f\n", symbol, prettyPrint(balance, int(decimals)))

}
