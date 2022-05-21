package main

import (
	"Vico1993/Wallet/service"
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	wallet := GetData()

	fmt.Printf("Number of Transactions: %s\n", strconv.Itoa(len(wallet.Transactions)))

	for _, transaction := range wallet.Transactions {
		price, err := service.GetAssetPrice(transaction.Asset)
		if err != nil {
			log.Fatalln(transaction.Asset, "---" , err)
		}

		fmt.Println("New price for ", transaction.Asset, "is: ", strconv.FormatFloat(price, 'g', -1, 64))

		// fmt.Printf(
		// 	"%s: You bough %s, at %s. Now it's %s\n",
		// 	transaction.Date,
		// 	transaction.Asset,
		// 	strconv.FormatFloat(transaction.AssetPrice, 'g', -1, 64),
		// 	strconv.FormatFloat(service.GetAssetPrice(transaction.Asset), 'g', -1, 64),
		// )
	}
}