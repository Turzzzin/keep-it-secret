package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password_hash"`
}

func getUserFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".kis")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(dir, "user.json"), nil
}

func userExists() bool {
	userFile, err := getUserFile()
	if err != nil {
		return false
	}
	_, err = os.Stat(userFile)
	return err == nil
}

func RegisterUser(username, password string) error {
	if userExists() {
		return errors.New("a user is already registered on this device")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Password: string(hashedPassword),
	}

	userFile, err := getUserFile()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(userFile, data, 0600); err != nil {
		return err
	}

	fmt.Println("✅ User registered successfully!")
	return nil
}

func AuthenticateUser(username, password string) error {
	userFile, err := getUserFile()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(userFile)
	if err != nil {
		return errors.New("no registered user found — please run `kis register` first")
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	if username != user.Username {
		return errors.New("invalid username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid password")
	}

	fmt.Println("🔓 Authentication successful!")
	return nil
}

func DeleteUser(username, password string) error {
	if err := AuthenticateUser(username, password); err != nil {
		return err
	}

	userFile, err := getUserFile()
	if err != nil {
		return err
	}

	if err := os.Remove(userFile); err != nil {
		return err
	}

	fmt.Println("✅ User deleted successfully!")
	return nil
}

func DeleteAllUsers(password string) error {
    // Since there is only one user, we can just get the username from the file
    userFile, err := getUserFile()
    if err != nil {
        return err
    }

    data, err := os.ReadFile(userFile)
    if err != nil {
        return errors.New("no registered user found — please run `kis register` first")
    }

    var user User
    if err := json.Unmarshal(data, &user); err != nil {
        return err
    }

	return DeleteUser(user.Username, password)
}