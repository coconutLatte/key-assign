package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

const timeFormat = "2006-01-02 15:04:05"

type cmdGenerate struct {
	flagUID           int
	flagValidDuration time.Duration
	flagSecret        string
}

func (c *cmdGenerate) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate key with expired time",
		Long:  "generate 10000 30d",
		RunE:  c.runE,
	}

	cmd.Flags().IntVar(&c.flagUID, "uid", -1, "user-id")
	cmd.Flags().DurationVar(&c.flagValidDuration, "valid-duration", 0, "valid duration [ 1s | 1h | 1m ]")
	cmd.Flags().StringVar(&c.flagSecret, "secret", "", "secret to encode token")

	return cmd
}

func (c *cmdGenerate) runE(_ *cobra.Command, _ []string) error {
	if c.flagUID == -1 {
		return fmt.Errorf("uid cannot be empty")
	}

	if c.flagValidDuration == 0 {
		return fmt.Errorf("valid-duration cannot be empty")
	}

	if c.flagSecret == "" {
		return fmt.Errorf("secret cannot be empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":        c.flagUID,
		"expired_at": time.Now().Add(c.flagValidDuration).Format(timeFormat),
	})

	str, err := token.SignedString(secret)
	if err != nil {
		return fmt.Errorf("sign string failed, %w", err)
	}

	fmt.Println(str)
	return nil
}
