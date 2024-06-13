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
	"log"
	"os"

	"github.com/blinklabs-io/bursa"
	"github.com/spf13/cobra"
)

func addressCommand() *cobra.Command {
	addressCommand := cobra.Command{
		Use:   "address",
		Short: "Address commands",
	}

	addressCommand.AddCommand(
		addressKeygenCommand(),
	)
	return &addressCommand
}

func addressKeygenCommand() *cobra.Command {
	walletCreateCommand := cobra.Command{
		Use:   "key-gen",
		Short: "Creates payment key pair from mnemonic",
		Run:   runAddressKeygenCommand,
	}

	return &walletCreateCommand
}

func runAddressKeygenCommand(cmd *cobra.Command, args []string) {
	paymentId, accountId := 0, 0 // as long as bursa has no state, will always be 0?
	rootKey, err := bursa.GetRootKeyFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("error creating root key: %s", err.Error())
	}
	// bursa should not force me to know about uint/uint32; Just accept int and
	// provide an error if out of range?
	accountKey := bursa.GetAccountKey(rootKey, uint(accountId))
	paymentKey := bursa.GetPaymentKey(accountKey, uint32(paymentId))
	paymentVKey := bursa.GetPaymentVKey(paymentKey)
	paymentSKey := bursa.GetPaymentSKey(paymentKey)

	err = writeKeyFile(paymentVKey, verificationKeyFile)
	if err != nil {
		log.Fatalf("error writing %s: %s", verificationKeyFile, err.Error())
	}
	err = writeKeyFile(paymentSKey, signingKeyFile)
	if err != nil {
		log.Fatalf("error writing %s: %s", signingKeyFile, err.Error())
	}
}

func writeKeyFile(keyFile bursa.KeyFile, filePath string) error {
	f, err := json.MarshalIndent(keyFile, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, f, 0600)
	if err != nil {
		return err
	}
	return nil
}
