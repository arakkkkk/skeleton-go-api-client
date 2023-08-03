package main

import (
  "go-api-client-sample/samplesdk/items"
)

func main() {
  items, err := items.New("aa", "token", nil)
  if err != nil {
		panic(err)
	}
	items.GetUserItems()
}
