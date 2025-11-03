package main

import (
	"encoding/json"
	"fmt"
	"keep-it-secret/internal/auth"
	"keep-it-secret/internal/crypto"
	"keep-it-secret/internal/storage"
	"keep-it-secret/internal/ui"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "kis",
		Short: "Keep it Secret — CLI tool to keep your secrets safe",
		Long:  "Keep it Secret (kis) is a command-line tool for securely storing secrets locally using AES encryption.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(
		newRegisterCommand(),
		newAddCommand(),
		newGetCommand(),
		newListCommand(),
		newDeleteCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRegisterCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Create a new user account",
		RunE: func(cmd *cobra.Command, args []string) error {
			username := ui.Prompt("Enter username: ")
			password := ui.PromptPassword("Enter master password: ")
			confirm := ui.PromptPassword("Confirm password: ")
			if password != confirm {
				return fmt.Errorf("passwords do not match")
			}
			return auth.RegisterUser(username, password)
		},
	}
}

func newAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add <secret-name>",
		Short: "Add a new structured secret (multiple key/value pairs)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			fmt.Printf("Creating secret: %s\n", name)
			fmt.Println("Enter your key/value pairs (press Enter on an empty key to finish):")

			data := make(map[string]string)

			for {
				key := ui.Prompt("  Key: ")
				if key == "" {
					break
				}
				value := ui.Prompt(fmt.Sprintf("  Value for '%s': ", key))
				data[key] = value
			}

			if len(data) == 0 {
				fmt.Println("❌ No key/value pairs provided. Aborting.")
				return nil
			}

			testJSON, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("invalid data: %w", err)
			}

			master := ui.PromptPassword("Enter master password: ")

			enc, err := crypto.Encrypt(testJSON, master)
			if err != nil {
				return err
			}

			store, err := storage.New()
			if err != nil {
				return err
			}

			secret := storage.Secret{
				Name:      name,
				Data:      data,
				Encrypted: enc,
			}

			if err := store.SaveSecret(name, secret.Encrypted); err != nil {
				return err
			}

			fmt.Println("✅ Secret saved successfully!")
			return nil
		},
	}
}

func newGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get <secret-name>",
		Short: "Retrieve and decrypt a stored secret",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			master := ui.PromptPassword("Enter master password: ")

			store, err := storage.New()
			if err != nil {
				return err
			}

			secret, err := store.GetSecret(name)
			if err != nil {
				return err
			}

			plaintext, err := crypto.Decrypt(secret.Encrypted, master)
			if err != nil {
				return err
			}

			var data map[string]string
			if err := json.Unmarshal(plaintext, &data); err == nil {
				fmt.Printf("🔓 %s:\n", name)
				keys := make([]string, 0, len(data))
				for k := range data {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				for _, k := range keys {
					fmt.Printf("  %s: %s\n", k, data[k])
				}
			} else {
				fmt.Printf("🔓 %s: %s\n", name, string(plaintext))
			}

			return nil
		},
	}
}

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all stored secrets",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("📜 Listing secrets...")

			store, err := storage.New()
			if err != nil {
				return err
			}

			names := store.ListSecrets()
			if len(names) == 0 {
				fmt.Println("  (no secrets found)")
				return nil
			}

			sort.Strings(names)

			for _, n := range names {
				fmt.Println(" -", n)
			}
			return nil
		},
	}
}

func newDeleteCommand() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a user or a secret",
	}

	deleteUserCmd := &cobra.Command{
		Use:   "user [username]",
		Short: "Delete a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			all, _ := cmd.Flags().GetBool("all")
			if all {
				password := ui.PromptPassword("Enter master password to delete all users: ")
				return auth.DeleteAllUsers(password)
			} else {
				if len(args) == 0 {
					return fmt.Errorf("username is required")
				}
				username := args[0]
				password := ui.PromptPassword("Enter master password to delete user: ")
				return auth.DeleteUser(username, password)
			}
		},
	}

	deleteUserCmd.Flags().Bool("all", false, "Delete all users")

	deleteSecretCmd := &cobra.Command{
		Use:   "secret [secret-name]",
		Short: "Delete a secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			all, _ := cmd.Flags().GetBool("all")
			store, err := storage.New()
			if err != nil {
				return err
			}

			if all {
				return store.ClearAll()
			} else {
				if len(args) == 0 {
					return fmt.Errorf("secret name is required")
				}
				name := args[0]
				return store.DeleteSecret(name)
			}
		},
	}

	deleteSecretCmd.Flags().Bool("all", false, "Delete all secrets")

	deleteCmd.AddCommand(deleteUserCmd, deleteSecretCmd)

	return deleteCmd
}