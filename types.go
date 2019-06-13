package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

type (
	base struct {
		OK bool
	}

	radiusSearch struct {
		base
		License string
	}
)

func main() {
	resp, err := resty.R().Get("https://creativecommons.tankerkoenig.de/json/list.php?lat=52.521&lng=13.438&rad=1.5&sort=dist&type=all&apikey=00000000-0000-0000-0000-000000000002")
	if err != nil {
		panic(err)
	}
	var s radiusSearch
	err = json.Unmarshal(resp.Body(), &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.License)
}
