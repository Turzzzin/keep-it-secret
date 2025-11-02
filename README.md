 ____  __.                     .___  __      _________                            __
|    |/ _|____   ____ ______   |   |/  |_   /   _____/ ____   ___________   _____/  |_
|      <_/ __ \_/ __ \\____ \  |   \   __\  \_____  \_/ __ \_/ ___\_  __ \_/ __ \   __\
|    |  \  ___/\  ___/|  |_> > |   ||  |    /        \  ___/\  \___|  | \/\  ___/|  |
|____|__ \___  >\___  >   __/  |___||__|   /_______  /\___  >\___  >__|    \___  >__|
        \/   \/     \/|__|                         \/     \/     \/            \/

A command-line interface (CLI) tool for securely storing secrets locally.

## Features

*   **User Authentication:** Create a local account with a username and password.
*   **Secure Secret Storage:** Encrypt and store your secrets on your local machine.
*   **Master Password:** A single master password is used to encrypt and decrypt all your secrets, ensuring that only you can access them.
*   **Cross-Platform:** Built with Go, `keep-it-secret` can be compiled for various operating systems.

## Architecture

The `keep-it-secret` CLI is built with a modular architecture, separating concerns into different packages.

*   **`cmd/keep-it-secret`**: The main application entry point. It handles command-line argument parsing and orchestrates calls to the other internal packages. We use the `cobra` library for creating a powerful and modern CLI application.
*   **`internal/auth`**: Handles user authentication. This includes creating new users, verifying credentials, and managing user sessions. Passwords are hashed before being stored.
*   **`internal/crypto`**: Manages all cryptographic operations. It uses AES-256-GCM for strong encryption of secrets. The master password is used to derive an encryption key.
*   **`internal/storage`**: Responsible for all interactions with the local filesystem. It stores user data and the encrypted secrets in a designated directory.
*   **`internal/ui`**: Handles all user-facing interactions, such as prompting for passwords and displaying formatted output.

## Folder Structure

```
/
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── cmd/
│   └── keep-it-secret/
│       └── main.go         # Main application entry point
├── internal/
│   ├── auth/
│   │   └── auth.go         # User authentication logic
│   ├── crypto/
│   │   └── crypto.go       # Encryption and decryption logic
│   ├── storage/
│   │   └── storage.go      # Local file storage logic
│   └── ui/
│       └── ui.go           # User interface and interaction logic
├── scripts/
│   └── install.sh          # Installation script
└── docs/
    └── architecture.md     # Detailed architecture documentation
```

## Installation

1.  **Prerequisites:** Make sure you have Go installed on your system.
2.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/keep-it-secret.git
    cd keep-it-secret
    ```
3.  **Run the installation script:**
    ```bash
    chmod +x scripts/install.sh
    ./scripts/install.sh
    ```
    The `install.sh` script will build the Go project and move the binary to a directory in your `PATH` (e.g., `/usr/local/bin`).

    **`scripts/install.sh`:**
    ```sh
    #!/bin/bash
    echo "Building keep-it-secret..."
    go build -o keep-it-secret ./cmd/keep-it-secret
    echo "Installing keep-it-secret to /usr/local/bin..."
    sudo mv keep-it-secret /usr/local/bin/
    echo "Installation complete!"
    ```

## Usage

### Create an account

```bash
kis register
```

### Add a new secret

```bash
kis add
```
You will be prompted for your master password and the secret you want to add.

### View a secret

```bash
kis get <secret-name>
```
You will be prompted for your master password to decrypt the secret.

### List all secrets

```bash
kis list
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License.
