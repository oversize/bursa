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

	"github.com/blinklabs-io/bursa"
	"github.com/spf13/cobra"
)

func keyCommand() *cobra.Command {
	keyCommand := cobra.Command{
		Use:   "key",
		Short: "Key commands",
	}

	keyCommand.AddCommand(
		keyCreateCommand(),
	)
	return &keyCommand
}

func keyCreateCommand() *cobra.Command {
	keyCreateCommand := cobra.Command{
		Use:   "create",
		Short: "Creates new root key from mnemonic",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Menemonic %s\n", mnemonic)
			rootkey, err := bursa.GetRootKeyFromMnemonic(mnemonic)
			if err != nil {
				log.Fatalf("create rootkey failed: %s", err.Error())
			}

			fmt.Println(rootkey)
		},
	}

	return &keyCreateCommand
}
