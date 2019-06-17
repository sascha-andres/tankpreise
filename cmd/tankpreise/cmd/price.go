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
	"github.com/spf13/viper"
	"livingit.de/code/tankpreise"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Get prices for selected gas stations",
	Long:  `Provide a list of prices for selected gas stations`,
	Run: func(cmd *cobra.Command, args []string) {
		gp, err := tankpreise.NewGasPrices()
		if err != nil {
			panic(err)
		}
		gp.SetLicense(viper.GetString("license"))
		stations := viper.GetStringSlice("price.station-id")
		if len(stations) == 0 {
			_, _ = fmt.Fprintln(os.Stderr, "you need to provide at least one station")
			os.Exit(1)
		}
		p, err := gp.PriceQuery(tankpreise.PricesRequest{
			IDs: stations,
		})
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
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
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.Flags().StringSliceP("station-id", "", nil, "provide id of station (may be provided multiple times)")
	_ = viper.BindPFlag("price.station-id", priceCmd.Flags().Lookup("station-id"))
}
