/*
Copyright © 2019 Sascha Andres <sascha.andres@outlook.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"livingit.de/code/tankpreise"
	"os"
	"text/tabwriter"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for gas stations",
	Long:  `Search for gas stations around a location.`,
	Run: func(cmd *cobra.Command, args []string) {
		gp, err := tankpreise.NewGasPrices()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		gp.SetLicense(viper.GetString("license"))
		s, err := gp.Search(tankpreise.SearchRequest{
			GasType:   viper.GetString("search.gas-type"),
			Latitude:  viper.GetFloat64("search.latitude"),
			Longitude: viper.GetFloat64("search.longitude"),
			Radius:    viper.GetFloat64("search.radius"),
			Sort:      viper.GetString("search.sort"),
		})
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		const padding = 3
		w := tabwriter.NewWriter(os.Stdout, 0, 2, padding, ' ', 0)
		_, _ = fmt.Fprintln(w, "ID\tOpen\tName\tStreet\tZipCode\tCity")
		for _, station := range s.Stations {
			_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t%t\t%s\t%s\t%d\t%s", station.ID, station.IsOpen, station.Name, fmt.Sprintf("%s %s", station.Street, station.HouseNumber), station.PostCode, station.Place))
		}
		err = w.Flush()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	searchCmd.Flags().Float64P("latitude", "", 52.521, "latitude of location")
	searchCmd.Flags().Float64P("longitude", "", 13.438, "longitude of location")
	searchCmd.Flags().Float64P("radius", "r", 10.0, "radius around location for search")
	searchCmd.Flags().StringP("gas-type", "g", "all", "type of gas you're looking for (diesel, e5, e10 or all)")
	searchCmd.Flags().StringP("sort", "s", "dist", "sort by (dist)ance or price")

	_ = viper.BindPFlag("search.latitude", searchCmd.Flags().Lookup("latitude"))
	_ = viper.BindPFlag("search.longitude", searchCmd.Flags().Lookup("longitude"))
	_ = viper.BindPFlag("search.radius", searchCmd.Flags().Lookup("radius"))
	_ = viper.BindPFlag("search.gas-type", searchCmd.Flags().Lookup("gas-type"))
	_ = viper.BindPFlag("search.sort", searchCmd.Flags().Lookup("sort"))

}
