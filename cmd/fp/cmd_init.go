package main

import (
	"os"

	"fmt"

	"bytes"

	"github.com/s-kirby/fastpass"
	"github.com/s-kirby/fastpass/keyfile"
	"github.com/pkg/errors"
)

func cmdInit() {
	fp := fastpass.New()
	if _, err := os.Stat(config.DB); err == nil {
		fail("db @ %v exists", config.DB)
	}

	var key [32]byte
	var err error

	if config.KeyFile == "" {
		pwd := fastpass.GetPassword()
		fmt.Printf("(confirm) ")
		cpwd := fastpass.GetPassword()

		if bytes.Compare(pwd[:], cpwd[:]) != 0 {
			fail("password mismatch")
		}

		key = pwd
	} else {
		if key, err = keyfile.Load(config.KeyFile); os.IsNotExist(errors.Cause(err)) {
			if key, err = keyfile.Create(config.KeyFile); err != nil {
				fail("failed to create key file @ %v: %v", config.KeyFile, err)
			}
		} else if err != nil {
			fail("unexpected error while loading %v: %v", config.KeyFile, err)
		}
		info("using key file @ %v", config.KeyFile)
	}
	fp.Key = key

	if err := fp.Create(config.DB); err != nil {
		fail("failed to create db: %v", err)
	}

	fp.Close()

	success("created db @ %v", config.DB)
}
