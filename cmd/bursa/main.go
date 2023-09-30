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

package main

import (
	"fmt"
	"os"
)

func main() {
	var subCommand string
	// Parse subcommand (default: "cli")
	if len(os.Args) < 2 {
		subCommand = "cli"
	} else {
		subCommand = os.Args[1]
	}

	switch subCommand {
	case "cli":
		cliRun()
	default:
		fmt.Printf("Unknown subcommand: %s\n", subCommand)
		os.Exit(1)
	}
}
