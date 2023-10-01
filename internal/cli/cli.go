// Copyright 2023 Blink Labs, LLC.
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

package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/blinklabs-io/bursa"
	"github.com/blinklabs-io/bursa/internal/config"
	"github.com/blinklabs-io/bursa/internal/logging"
)

func Run() {
	fs := flag.NewFlagSet("cli", flag.ExitOnError)
	flagOutput := fs.String("output", "", "output directory for files, otherwise uses STDOUT")
	if len(os.Args) >= 2 {
		_ = fs.Parse(os.Args[2:]) // ignore parse errors
	}

	cfg := config.GetConfig()
	logger := logging.GetLogger()
	// Load mnemonic
	var err error
	mnemonic := cfg.Mnemonic
	if mnemonic == "" {
		mnemonic, err = bursa.NewMnemonic()
		if err != nil {
			logger.Fatalf("failed to load mnemonic: %s", err)
		}
	}
	w, err := bursa.NewDefaultWallet(mnemonic)
	if err != nil {
		logger.Fatalf("failed to initialize wallet: %s", err)
	}

	logger.Infof("Loaded mnemonic and generated address...")

	if *flagOutput == "" {
		fmt.Printf("MNEMONIC=%s\n", w.Mnemonic)
		fmt.Printf("PAYMENT_ADDRESS=%s\n", w.PaymentAddress)
		fmt.Printf("STAKE_ADDRESS=%s\n", w.StakeAddress)

		fmt.Printf("payment.vkey=%s\n", bursa.GetKeyFile(w.PaymentVKey))
		fmt.Printf("payment.skey=%s\n", bursa.GetKeyFile(w.PaymentSKey))
		fmt.Printf("stake.vkey=%s\n", bursa.GetKeyFile(w.StakeVKey))
		fmt.Printf("stake.skey=%s\n", bursa.GetKeyFile(w.StakeSKey))
	} // TODO: output to files
}