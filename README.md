# flare-claim-rewards
This is a tool to claim rewards from the Flare Network FTSO V2.

## Run script
Set up .env file with FTSO_V1_WALLET and FTSO_V2_IDENTITY.
```
go run cmd/rewardmanager/main.go ./data/reward_claims.json
```

## Set Up Development Environment

### Prerequisites
- Go 1.19 or higher
- github.com/ethereum/go-ethereum
- github.com/stretchr/testify

### Installing Dependencies

```
go mod download
```

## Running Tests

To run all tests in the parser package:
```
go test ./pkg/parser
```