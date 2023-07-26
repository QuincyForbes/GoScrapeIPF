package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)



type Metadata struct {
	Cid         string `json:"cid"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Cid struct {
	Cid string `json:"cid"`
}



func fetchMetadataFromCID(cid string) (*Metadata, error) {
	ipfsGatewayURL := "https://blockpartyplatform.mypinata.cloud/ipfs/"
	finalGateway := ipfsGatewayURL + cid

	const maxAttempts = 5

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := http.Get(finalGateway)
		if err != nil {
			fmt.Printf("error retrieving data for CID %s on attempt %d: %s\n", cid, attempt, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("non-200 status code for CID %s on attempt %d: %d\n", cid, attempt, resp.StatusCode)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error reading response body for CID %s on attempt %d: %s\n", cid, attempt, err)
			continue
		}

		var metadata Metadata
		if err := json.Unmarshal(body, &metadata); err != nil {
			fmt.Printf("failed to decode JSON for CID %s on attempt %d: %s\n", cid, attempt, err)
			continue
		}

		metadata.Cid = cid
		return &metadata, nil
	}


	return nil, nil

}


	func readCIDsFromCSV(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	var cids []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		cids = append(cids, record[0])
	}
	return cids, nil
}


func putItemToDynamoDB(item *Metadata) error {
	
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	
	svc := dynamodb.NewFromConfig(cfg)

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshal Record, %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("MetadataTable"),
		Item:      av,
	}

	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to put Record to DynamoDB, %w", err)
	}

	return nil
}

	

func main() {
	cids, err := readCIDsFromCSV("ipfs_cids.csv")
	if err != nil {
		fmt.Printf("Error reading CIDs from CSV: %s\n", err)
		return
	}

	for _, cid := range cids {
		metadata, err := fetchMetadataFromCID(cid)
		if err != nil {
			fmt.Printf("Error fetching metadata for CID %s: %s\n", cid, err)
			continue
		}
		
		err = putItemToDynamoDB(metadata)
		if err != nil {
			fmt.Printf("Error saving metadata for CID %s to DynamoDB: %s\n", cid, err)
			continue
		}
	}
}
