/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ktr0731/emvutil/format"
	"github.com/spf13/cobra"
	"go.mercari.io/go-emv-code/mpm"
	"go.mercari.io/go-emv-code/mpm/jpqr"
	"go.mercari.io/go-emv-code/tlv"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode input as a JPQR payload",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var in io.Reader
		if len(args) != 0 {
			in = strings.NewReader(strings.Join(args, "\n"))
		} else {
			in = os.Stdin
		}

		var fmter format.Formatter
		switch {
		case *json:
			fmter = format.NewJSON(os.Stdout)
		default:
			fmter = format.NewPP(os.Stdout)
		}

		var hasErr bool
		s := bufio.NewScanner(in)
		for s.Scan() {
			if *decodeAsID {
				code := &mpm.Code{
					MerchantAccountInformation: []tlv.TLV{
						{
							Tag:    "26",
							Length: "68",
							Value:  s.Text(),
						},
					},
				}
				id, err := jpqr.ParseID(code)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to decode id '%s'", s.Text())
					hasErr = true
					continue
				}
				fmter.Format(id)
			} else {
				code, err := jpqr.Decode(s.Bytes())
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to decode payload '%s'", s.Text())
					hasErr = true
					continue
				}
				fmter.Format(code)
			}
		}
		if hasErr {
			os.Exit(1)
		}
	},
}

var (
	decodeAsID *bool
	json       *bool
)

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	decodeAsID = decodeCmd.Flags().Bool("id", false, "decode input as an JPQR ID")
	json = decodeCmd.Flags().Bool("json", false, "format as a JSON text")
}
