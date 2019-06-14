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
	p, err := gp.PriceQuery(tankpreise.PricesRequest{
		IDs: []string{"4429a7d9-fb2d-4c29-8cfe-2ca90323f9f8", "446bdcf5-9f75-47fc-9cfa-2c3d6fda1c3b", "60c0eefa-d2a8-4f5c-82cc-b5244ecae955", "44444444-4444-4444-4444-444444444444"},
	})
	if err != nil {
		panic(err)
	}
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 2, padding, ' ', 0)
	_, _ = fmt.Fprintln(w, "ID\tStatus\tE5\tE10\tDiesel")
	for key, data := range p.Prices {
		e5 := 0.0
		if data.HasE5() {
			e5, _ = data.GetE5()
		}
		e10 := 0.0
		if data.HasE10() {
			e10, _ = data.GetE10()
		}
		diesel := 0.0
		if data.HasDiesel() {
			diesel, _ = data.GetDiesel()
		}
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%f\t%f\t%f", key, data.Status, e5, e10, diesel))
	}
	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func search() {
	gp, err := tankpreise.NewGasPrices()
	if err != nil {
		panic(err)
	}
	s, err := gp.Search(tankpreise.SearchRequest{
		GasType:   "e10",
		Latitude:  52.521,
		Longitude: 13.438,
		Radius:    10.0,
		Sort:      "dist",
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
