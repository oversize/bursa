// Copyright 2023 Blink Labs Software
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/blinklabs-io/bursa/internal/api"
	"github.com/blinklabs-io/bursa/internal/config"
	"github.com/blinklabs-io/bursa/internal/logging"
	"github.com/spf13/cobra"
)

func apiCommand() *cobra.Command {
	apiCommand := cobra.Command{
		Use:   "api",
		Short: "Runs the api",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.GetConfig()
			// Start API listener
			logger := logging.GetLogger()
			// Start API listener
			logger.Infof(
				"starting API listener on %s:%d",
				cfg.Api.ListenAddress,
				cfg.Api.ListenPort,
			)
			if err := api.Start(cfg); err != nil {
				logger.Fatalf("failed to start API: %s", err)
			}

			// Wait forever
			select {}
		},
	}
	return &apiCommand
}
