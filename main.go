package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/denis-zakharov/goswagger-client/client"
	"github.com/denis-zakharov/goswagger-client/client/products"
)

func main() {
	// c := client.Default
	transportConfig := client.DefaultTransportConfig().WithHost("localhost:3000")
	c := client.NewHTTPClientWithConfig(nil, transportConfig)
	params := products.NewListProductsParams()
	resp, err := c.Products.ListProducts(params)
	if err != nil {
		log.Fatalf("List products error: %v", err)
	}
	for _, v := range resp.Payload {
		d, err := json.Marshal(v)
		if err != nil {
			log.Printf("[ERROR] serializing to json %+v\n", v)
		}
		fmt.Println(string(d))
	}
}
