# IPFS Metadata Fetcher

This project fetches metadata from IPFS CIDs and saves them to an AWS DynamoDB table. Specifically, the program reads a list of CIDs from a CSV file, retrieves their metadata from an IPFS gateway, and subsequently saves each metadata record to a DynamoDB table.

## Features

- **Fetch Metadata**: Get the metadata of an IPFS CID from the specified gateway.
- **Retries**: The program attempts to fetch metadata up to five times before giving up on a particular CID.
- **Save to DynamoDB**: After fetching, the metadata is saved into a DynamoDB table named "MetadataTable".

## Prerequisites

- **Go**: This is a Go project, and you would need Go installed to run or build it.
- **AWS CLI & AWS SDK**: The AWS CLI needs to be set up with appropriate permissions. The program also utilizes the AWS SDK for Go v2.
- **IPFS Gateway**: This project specifically uses `https://blockpartyplatform.mypinata.cloud/ipfs/` as its default IPFS gateway. You can modify it in the code if you use a different gateway.
- **CSV with CIDs**: A CSV file named "ipfs_cids.csv" containing a list of CIDs to be processed.

## Usage

1. Ensure you have a file named `ipfs_cids.csv` in the same directory as the program. This CSV should contain CIDs, one per line.
2. Run the program:
   `go run main.go`

Upon execution, the program will:

1. Read the CIDs from the CSV.
2. For each CID, fetch its metadata from the IPFS gateway.
3. Save the metadata to the DynamoDB table.
