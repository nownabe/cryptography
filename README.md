cryptography
============

cryptography is a command line tool to encrypt/decrypt files.

# Install
```bash
go get -u github.com/nownabe/cryptography
```

# Encrypt
The length of encryption_key must be 16 or 32 bytes.

```bash
cryptography enc -in input_path -out output_path -key encryption_key
```

# Decrypt
```bash
cryptography dec -in input_path -out output_path -key encryption_key
```
