# Toy Blockchain

## Overview

This is a simple blockchain implementation written in Go. It demonstrates the core concepts of blockchain technology, including block creation, chain validation, and the ability to detect tampering through hash verification.

## Project Structure

The project is organized into the following packages:

- **blockchain**: Contains the main blockchain logic, including block creation, chain management, and validation.
- **files_oper**: Handles file operations, specifically reading data from files to be stored in blocks.
- **main.go**: The entry point that demonstrates how to use the blockchain.

## Features

- **Block Creation**: Create blocks with timestamps, values, and cryptographic hashes.
- **Genesis Block**: Automatically creates the first block in the chain when the first data is added.
- **Chain Validation**: Verify the integrity of the entire blockchain by checking all hashes and previous hash references.
- **Hash Computation**: Uses SHA-256 to compute secure block hashes.
- **Chain Display**: Print the contents of all blocks in the chain for inspection.

## How It Works

### Creating a Blockchain

The blockchain is implemented as a linked list of blocks. Each block contains:

- An index (position in the chain)
- A timestamp
- Data value
- A hash (computed from the block's contents)
- A reference to the previous block's hash

### Adding Blocks

New blocks are added to the end of the chain. Each new block automatically references the hash of the previous block, creating an unbreakable link between them.

### Validating the Chain

The blockchain can be validated by checking that each block's hash is correct and matches the previous hash reference in the next block. Any tampering will be detected because the hashes will no longer align.

## Requirements

- Go 1.24.1 or later

## Usage

1. Prepare text files with data to be stored in blocks (for example, `testG.txt`, `test0.txt`, `test1.txt`, `test2.txt`).
2. Run the program:

```bash
go run main.go
```

The program will:

- Read data from the files
- Create a blockchain with multiple blocks
- Demonstrate chain validation
- Print the complete blockchain

## Example Output

The program displays each block's information including its index, timestamp, value, hash, and previous hash. It also shows whether the chain is valid after any modifications.

## Notes

This is a simplified educational implementation of blockchain technology. It is not intended for production use. Real blockchain systems use more advanced cryptography, consensus mechanisms, and security measures.
