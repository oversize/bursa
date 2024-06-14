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
	"os"

	"github.com/blinklabs-io/bursa"
	"github.com/spf13/cobra"
)

var (
	// as long as bursa has no state, will always be 0?
	paymentId, accountId, stakeId = 0, 0, 0
)

func keyCommand() *cobra.Command {
	keyCommand := cobra.Command{
		Use:   "key",
		Short: "Key commands",
	}

	keyPaymentCommand := cobra.Command{
		Use:   "payment",
		Short: "Create new payment key (from mnemonic)",
		RunE:  createPaymentKey,
	}
	keyStakingCommand := cobra.Command{
		Use:   "staking",
		Short: "Create new staking key (from mnemonic)",
		RunE:  createStakingKey,
	}
	keyCommand.AddCommand(
		&keyPaymentCommand,
		&keyStakingCommand,
	)
	return &keyCommand
}

func createPaymentKey(cmd *cobra.Command, args []string) error {
	rootKey, err := bursa.GetRootKeyFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	accountKey := bursa.GetAccountKey(rootKey, uint(accountId))
	paymentKey := bursa.GetPaymentKey(accountKey, uint32(paymentId))

	if verificationKeyFile != "" {
		paymentVKey := bursa.GetPaymentVKey(paymentKey)
		err = writeKeyFile(paymentVKey, verificationKeyFile)
		if err != nil {
			return err
		}
	}

	if signingKeyFile != "" {
		paymentSKey := bursa.GetPaymentSKey(paymentKey)
		err = writeKeyFile(paymentSKey, signingKeyFile)
		if err != nil {
			return err
		}
	}

	if extendedSigningKeyFile != "" {
		paymentExtendedSKey := bursa.GetPaymentExtendedSKey(paymentKey)
		err = writeKeyFile(paymentExtendedSKey, extendedSigningKeyFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func createStakingKey(cmd *cobra.Command, args []string) error {
	rootKey, err := bursa.GetRootKeyFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	accountKey := bursa.GetAccountKey(rootKey, uint(accountId))
	stakeKey := bursa.GetStakeKey(accountKey, uint32(stakeId))

	if verificationKeyFile != "" {
		stakingVKey := bursa.GetStakeVKey(stakeKey)
		err = writeKeyFile(stakingVKey, verificationKeyFile)
		if err != nil {
			return err
		}
	}

	if signingKeyFile != "" {
		stakingSKey := bursa.GetStakeSKey(stakeKey)
		err = writeKeyFile(stakingSKey, signingKeyFile)
		if err != nil {
			return err
		}
	}

	if extendedSigningKeyFile != "" {
		stakingExtendedSKey := bursa.GetStakeExtendedSKey(stakeKey)
		err = writeKeyFile(stakingExtendedSKey, extendedSigningKeyFile)
		if err != nil {
			return err
		}
	}

	return nil
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
