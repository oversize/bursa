// Copyright 2024 Blink Labs Software
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
	"fmt"
	"log"
	"slices"

	"github.com/spf13/cobra"

	"github.com/blinklabs-io/bursa"
)

var (
	size int
)

func mnemonicCommand() *cobra.Command {
	mnemonicCommand := cobra.Command{
		Use:   "mnemonic",
		Short: "Command to create mnemonic (recovery-phrase)",
	}

	mnemonicCommand.AddCommand(
		mnemonicCreateCommand(),
	)
	return &mnemonicCommand
}

func mnemonicCreateCommand() *cobra.Command {
	valid_sizes := []int{9, 12, 15, 18, 21, 24}
	mnemonicCreateCommand := cobra.Command{
		Use:   "create",
		Short: "Creates a new mnemonic",
		Run: func(cmd *cobra.Command, args []string) {
			if !slices.Contains(valid_sizes, size) {
				log.Fatalf("Wrong size %s", size)
			}
			mnemonic, err := bursa.NewMnemonic()
			if err != nil {
				log.Fatalf("failed to load mnemonic: %s", err)
			}
			fmt.Println(mnemonic)
		},
	}

	mnemonicCreateCommand.Flags().IntVar(&size, "size", 24, "Specify length of mnemonic")

	return &mnemonicCreateCommand
}
