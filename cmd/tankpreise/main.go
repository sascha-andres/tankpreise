package main

import (
	"fmt"
	"livingit.de/code/tankpreise"
	"os"
	"text/tabwriter"
)

func main() {
	gp, err := tankpreise.NewGasPrices()
	if err != nil {
		panic(err)
	}
	s, err := gp.Search(tankpreise.SearchRequest{
		GasType:   "e10",
		Latitude:  52.521,
		Longitude: 13.438,
		Radius:    10.0,
	})
	if err != nil {
		panic(err)
	}
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 2, padding, ' ', 0)
	_, _ = fmt.Fprintln(w, "Open\tName\tStreet\tZipCode\tCity")
	for _, station := range s.Stations {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%t\t%s\t%s\t%d\t%s", station.IsOpen, station.Name, fmt.Sprintf("%s %s", station.Street, station.HouseNumber), station.PostCode, station.Place))
	}
	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
