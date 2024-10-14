# Letter Encryption Application

## Overview
This application implements a simple letter encryption algorithm in Go. It utilizes a basic rotation cipher to encrypt and decrypt text based on a specified key.

## Features
- Encrypts plain text using a rotational cipher.
- Decrypts previously encrypted text.
- Easy to use with a straightforward command line interface.

## How It Works
The encryption algorithm works by rotating the letters in the alphabet based on a specified key. For example, if the key is 2, 'A' becomes 'C', 'B' becomes 'D', and so on.

## Installation
1. Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. Clone the repository:
   ```bash
   git clone <repository-url>

Navigate to the project directory:
    cd <project-directory>

## Usage
To run the application, use the following command in your terminal:
    go run main.go

## Example
The following example demonstrates the encryption and decryption process:
    plainText := "ABCDE"
    fmt.Println("Plain Text:", plainText)

    encrypted := encrypt(2, plainText)
    fmt.Println("Encrypted:", encrypted)

    decrypted := decrypt(2, encrypted)
    fmt.Println("Decrypted:", decrypted)

## Command Line Arguments
Key: The number of positions each letter will be shifted in the alphabet.
Plain Text: The text you want to encrypt.

## Functions

### hashLetterFn
Description: Hashes the letter based on the key.
Parameters:
key: An integer representing the rotation amount.
letter: The string of letters to be hashed.
Returns: A new string with the letters rearranged.

### encrypt
Description: Encrypts the given plain text using the provided key.
Parameters:
key: The rotation amount.
plainText: The text to encrypt.
Returns: The encrypted string.

### decrypt
Description: Decrypts the given encrypted text using the provided key.
Parameters:
key: The rotation amount.
encrypttedText: The text to decrypt.
Returns: The decrypted string.

## Conclusion
This application provides a basic implementation of a letter encryption system. It is a great starting point for understanding how simple ciphers work.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
