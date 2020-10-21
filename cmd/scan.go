/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
    "os"
    "strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	portscanner "github.com/xzased/scanner/pkg/portscanner"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [host] [port]",
	Short: "Scan a host/port",
	Long: `Scan a host at a given port and try to get the data information from the service if any`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
	    port, err := strconv.Atoi(args[1])
	    if err != nil {
	        log.Error().Msg("Port must be an integer")
	        os.Exit(1)
	    }
        sr := portscanner.ScanRequest{
            Host: args[0],
            Port: port,
        }
        res := portscanner.ScanPort(sr)
        if res.Open {
            log.Info().
                Str("state", "open").
                Str("host", res.Host).
                Int("port", res.Port).
                Str("data", res.Data).
                Msg("found service")
        } else if res.Err != nil {
            log.Debug().
                Str("state", "error").
                Str("host", res.Host).
                Int("port", res.Port).
                Err(res.Err).
                Msg("error scanning")
        }
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
