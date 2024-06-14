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
	"encoding/json"
	"fmt"
	"os"

	"github.com/blinklabs-io/bursa"
	"github.com/spf13/cobra"
)

func addressCommand() *cobra.Command {
	addressCommand := cobra.Command{
		Use:   "address",
		Short: "Address commands",
	}
	addressCreateCommand := cobra.Command{
		Use:   "create",
		Short: "Create address from vkey (payment & staking)",
		RunE:  createAddress,
	}

	addressCommand.AddCommand(
		&addressCreateCommand,
	)
	return &addressCommand
}

func createAddress(cmd *cobra.Command, args []string) error {
	pKey, err := readKeyFile(paymentVerificationKeyFile)
	if err != nil {
		return err
	}
	sKey, err := readKeyFile(stakeVerificationKeyFile)
	if err != nil {
		return err
	}
	fmt.Print(pKey)
	fmt.Print(sKey)

	return nil
}

func readKeyFile(keyFile string) (*bursa.KeyFile, error) {
	k, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s ", keyFile)
	}
	key := bursa.KeyFile{}
	err = json.Unmarshal(k, &key)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal %s", keyFile)
	}
	return &key, nil
}

/*
func runPaymentKeygenCommand(cmd *cobra.Command, args []string) error {

		--payment-verification-key-file payment.vkey \
		--out-file payment.addr \
		--testnet-magic 2


	return nil
}
*/
