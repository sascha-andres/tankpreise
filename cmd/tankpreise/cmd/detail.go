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

// detailCmd represents the detail command
var detailCmd = &cobra.Command{
	Use:   "detail",
	Short: "print details about a station",
	Long: `Print out detailed information for a gas station. Includes:

- opening times
- prices
- address`,
	Run: func(cmd *cobra.Command, args []string) {
		gp, err := tankpreise.NewGasPrices()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		gp.SetLicense(viper.GetString("license"))
		details, err := gp.Detail(tankpreise.DetailRequest{
			ID: viper.GetString("detail.id"),
		})
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
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
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(detailCmd)
	detailCmd.Flags().StringP("station-id", "i", "", "provide ID of station")
	_ = detailCmd.MarkFlagRequired("station-id")
	_ = viper.BindPFlag("detail.id", detailCmd.Flags().Lookup("station-id"))
}
