package main

import (
	"fmt"
	"os"

	"fwdev.se/strongPasswordGenerator/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	length    int
	uppercase bool
	lowercase bool
	numbers   bool
	symbols   bool
	hex       bool
	base64    bool
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "pwgen",
		Short: "Generate secure passwords",
		RunE: func(cmd *cobra.Command, args []string) error {
			if hex {
				result, err := generator.GenerateHex(length)
				if err != nil {
					return err
				}
				fmt.Println(result)
				return nil
			}

			if base64 {
				result, err := generator.GenerateBase64(length)
				if err != nil {
					return err
				}
				fmt.Println(result)
				return nil
			}

			opts := generator.Options{
				Length:    length,
				Uppercase: uppercase,
				Lowercase: lowercase,
				Numbers:   numbers,
				Symbols:   symbols,
			}

			result, err := generator.GeneratePassword(opts)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Password length")
	rootCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "Include uppercase letters")
	rootCmd.Flags().BoolVarP(&lowercase, "lowercase", "w", false, "Include lowercase letters")
	rootCmd.Flags().BoolVarP(&numbers, "numbers", "n", false, "Include numbers")
	rootCmd.Flags().BoolVarP(&symbols, "symbols", "s", false, "Include symbols")
	rootCmd.Flags().BoolVar(&hex, "hex", false, "Generate hex string")
	rootCmd.Flags().BoolVar(&base64, "base64", false, "Generate base64 string")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
