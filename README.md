<pre>
    ____  __.                    __    .__  __      _________                            __
   |    |/ _|____   ____ _______/  |_  |__|/  |_   /   _____/ ____   ___________   _____/  |_
   |      <_/ __ \_/ __ \\____ \   __\ |  \   __\  \_____  \_/ __ \_/ ___\_  __ \_/ __ \   __\
   |    |  \  ___/\  ___/|  |_> >  |   |  ||  |    /        \  ___/\  \___|  | \/\  ___/|  |
   |____|__ \___  >\___  >   __/|__|   |__||__|   /_______  /\___  >\___  >__|    \___  >__|
           \/   \/     \/|__|                             \/     \/     \/            \/
</pre>

A command-line interface (CLI) tool for securely storing secrets locally.

`keep-it-secret` is a simple and powerful tool for managing your secrets directly from the terminal. It uses strong AES-256-GCM encryption to ensure that your sensitive information is always protected. With a single master password, you have complete control over your secrets, which are stored securely on your local machine.

## Installation

You can easily install `keep-it-secret` by running the following command in your terminal:

```bash
curl -sSL https://raw.githubusercontent.com/turzzzin/keep-it-secret/main/scripts/install.sh | bash
```

This script will download the latest release for your operating system and architecture, and install it in `/usr/local/bin`.

## Usage

### Register a new user

To start using `keep-it-secret`, you first need to create a user account. This account is local to your machine.

```bash
kis register
```

You will be prompted to enter a username and a master password. The master password is the key to all your secrets, so make sure to choose a strong one and keep it safe.

### Add a new secret

To add a new secret, use the `add` command. You can store multiple key-value pairs under a single secret name.

```bash
kis add <secret-name>
```

You will be prompted to enter your key-value pairs and your master password.

### Retrieve a secret

To retrieve a secret, use the `get` command.

```bash
kis get <secret-name>
```

You will be prompted for your master password to decrypt and display the secret.

### List all secrets

To see a list of all the secrets you have stored, use the `list` command.

```bash
kis list
```

### Delete a secret or user

You can delete a secret or a user with the `delete` command.

```bash
# Delete a secret
kis delete secret <secret-name>

# Delete all secrets
kis delete secret --all

# Delete a user
kis delete user <username>

# Delete all users
kis delete user --all
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License.