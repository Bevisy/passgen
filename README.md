# passgen
A Simple and Secure  Password Generator.

## Install
### brew
```sh
brew tap Bevisy/passgen
brew install passgen

passgen --help
```

### source
```sh
git clone git@github.com:Bevisy/passgen.git
make install

passgen --help
```

## Usage
```sh
❯ passgen --help
A simple and secure password generator

Usage:
  passgen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  random      Generate a random password
  uuid        Generate a UUID

Flags:
  -h, --help   help for passgen

Use "passgen [command] --help" for more information about a command.
```

## Example
```sh
# Generate random password
❯ passgen random
LDI5rOe8i541

# Generate a password of specified length
❯ passgen random -l 20
aNdMI7kgb3aJPB65MFTd

# Generate UUID
❯ passgen uuid
4129ac2653690082a488ce3eaf04afa6

# Format output
❯ passgen uuid -f
fe22239b-1983-8914-d065-d4093f72c327
```