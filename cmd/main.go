package main

import (
	"fmt"
	"livingit.de/code/tankpreise"
	"os"
	"text/tabwriter"
)

func main() {

}

func details() {
	gp, err := tankpreise.NewGasPrices()
	if err != nil {
		panic(err)
	}
	details, err := gp.Detail(tankpreise.DetailRequest{
		ID: "24a381e3-0d72-416d-bfd8-b2f65f6e5802",
	})
	if err != nil {
		panic(err)
	}
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 2, padding, ' ', 0)
	_, _ = fmt.Fprintln(w, fmt.Sprintf("ID\t%s", details.Station.ID))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Name\t%s", details.Station.Name))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Street\t%s %s", details.Station.Street, details.Station.HouseNumber))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("City\t%d %s", details.Station.PostCode, details.Station.Place))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("State\t%s", details.Station.State))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Brand\t%s", details.Station.Brand))
	_, _ = fmt.Fprintln(w, "Opening times\t")
	for _, val := range details.Station.OpeningTimes {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("\t%s (%s - %s)", val.Text, val.Start, val.End))
	}
	for _, val := range details.Station.Overrides {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("\t%s", val))
	}
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Is open\t%t", details.Station.IsOpen))
	err = w.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gasprices() {
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
