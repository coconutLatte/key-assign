package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

// TODO read it from config file, not set in code
var secret = []byte("secret")

type cmdDecode struct {
	flagSecret string
}

func (c *cmdDecode) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode <token>",
		Short: "decode key to check",
		RunE:  c.runE,
	}

	cmd.Flags().StringVar(&c.flagSecret, "secret", "", "secret for decode")

	return cmd
}

func (c *cmdDecode) runE(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expect 1 args but got %d", len(args))
	}

	token := args[0]
	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	jwtT, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return fmt.Errorf("parse jwt failed, %w", err)
	}
	claim, ok := jwtT.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("token invalid: cannot convert to jwt.MapClaims")
	}
	if !jwtT.Valid {
		return fmt.Errorf("token invalid")
	}

	d, err := json.Marshal(claim)
	if err != nil {
		return fmt.Errorf("marshal to []byte failed, %w", err)
	}

	fmt.Println(string(d))

	return nil
}
