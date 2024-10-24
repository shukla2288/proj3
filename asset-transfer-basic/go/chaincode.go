package main

import (
    "fmt"
    "encoding/json"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AssetContract struct {
    contractapi.Contract
}

type Asset struct {
    ID             string `json:"id"`
    Owner          string `json:"owner"`
    Description    string `json:"description"`
    Value          int    `json:"value"`
}

// CreateAsset adds a new asset to the ledger
func (c *AssetContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, owner string, description string, value int) error {
    asset := Asset{
        ID:          id,
        Owner:       owner,
        Description: description,
        Value:       value,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(id, assetJSON)
}

// TransferAsset changes the ownership of an existing asset
func (c *AssetContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil || assetJSON == nil {
        return fmt.Errorf("Asset not found")
    }
    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return err
    }
    asset.Owner = newOwner
    assetJSON, err = json.Marshal(asset)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(id, assetJSON)
}

// UpdateAsset updates the description and value of an existing asset
func (c *AssetContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, newDescription string, newValue int) error {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil || assetJSON == nil {
        return fmt.Errorf("Asset not found")
    }
    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return err
    }
    asset.Description = newDescription
    asset.Value = newValue
    assetJSON, err = json.Marshal(asset)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(id, assetJSON)
}

// QueryAsset retrieves an asset by its ID
func (c *AssetContract) QueryAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil || assetJSON == nil {
        return nil, fmt.Errorf("Asset not found")
    }
    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }
    return &asset, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(AssetContract))
    if err != nil {
        fmt.Printf("Error creating chaincode: %s", err.Error())
    }
    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting chaincode: %s", err.Error())
    }
}
