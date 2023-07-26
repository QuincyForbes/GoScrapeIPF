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


## Installation

1. Clone the repository to your local machine.
3. Run `go mod download` to download the project dependencies.
4. Set up your AWS credentials by running `aws configure`.
5. Create a DynamoDB table or update the code to support yours.


## Usage
Run the program:
   ```go run main.go```


## Results
![image](https://github.com/QuincyForbes/GoScrapeIPF/assets/74159902/e4428f02-6df9-49d6-a0ff-cac3a0e87cb8)
## Contributing

If you'd like to contribute to this project, please fork the repository and submit a pull request.


