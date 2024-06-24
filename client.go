package main

import (
	"context"
	"io"
	"net/http"

	"github.com/huangx8/oapi-codegen-share/example/api"
)

const (
	url = "http://localhost:8888"
)

func printResponse(response *http.Response) {
	if b, err := io.ReadAll(response.Body); err == nil {
		print(string(b))
	}
}

func main() {
	ctx := context.Background()
	httpClient := http.Client{}

	c, _ := api.NewClient(url, api.WithHTTPClient(&httpClient))

	resp, _ := c.ListCards(ctx, nil)
	printResponse(resp)

	cwr, _ := api.NewClientWithResponses(url, api.WithHTTPClient(&httpClient))
	addCardResponse, _ := cwr.AddCardWithResponse(ctx, api.AddCardJSONRequestBody{Owner: "hahaha"})
	// not handling 201
	println(addCardResponse.StatusCode())

	resp, _ = c.ListCards(ctx, nil)
	cardsResponse, _ := api.ParseListCardsResponse(resp)
	for _, card := range *cardsResponse.JSON200 {
		println(card.Id + ": " + card.Owner)
	}
}
