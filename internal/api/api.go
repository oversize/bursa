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

package api

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/blinklabs-io/bursa"
	"github.com/blinklabs-io/bursa/internal/config"
	"github.com/blinklabs-io/bursa/internal/logging"
)

// @title bursa
// @version v0
// @description Programmable Cardano Wallet API
// @Schemes http
// @BasePath /

// @contact.name Blink Labs
// @contact.url https://blinklabs.io
// @contact.email support@blinklabs.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func Start(cfg *config.Config) error {
	// Disable gin debug and color output
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	// Configure API router
	router := gin.New()
	// Catch panics and return a 500
	router.Use(gin.Recovery())
	// Access logging
	accessLogger := logging.GetAccessLogger()
	router.Use(ginzap.GinzapWithConfig(accessLogger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{},
	}))
	router.Use(ginzap.RecoveryWithZap(accessLogger, true))

	// Create a healthcheck
	router.GET("/healthcheck", handleHealthcheck)
	// Configure API routes
	router.GET("/api/wallet/create", handleWalletCreate)

	// Start API listener
	err := router.Run(fmt.Sprintf("%s:%d",
		cfg.Api.ListenAddress,
		cfg.Api.ListenPort,
	))
	return err
}

func handleHealthcheck(c *gin.Context) {
	c.JSON(200, gin.H{"healthy": true})
}

// handleCreateWallet godoc
//
// @Summary CreateWallet
// @Description Create a wallet and return details
// @Produce json
// @Success 200 bursa.Wallet string "Ok"
// @Router /api/wallet/create [get]
func handleWalletCreate(c *gin.Context) {
	logger := logging.GetLogger()

	mnemonic, err := bursa.NewMnemonic()
	if err != nil {
		logger.Errorf("failed to load mnemonic: %s", err)
		c.JSON(500, fmt.Sprintf("failed to load mnemonic: %s", err))
		return
	}

	w, err := bursa.NewDefaultWallet(mnemonic)
	if err != nil {
		logger.Errorf("failed to initialize wallet: %s", err)
		c.JSON(500, fmt.Sprintf("failed to initialize wallet: %s", err))
		return
	}
	c.JSON(200, w)
}