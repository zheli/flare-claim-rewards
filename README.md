# flare-claim-rewards
This is a tool to claim rewards from the Flare Network FTSO V2.

## Run script
1. Clone Flare Rewards repository https://github.com/flare-foundation/fsp-rewards/tree/main
```
git clone https://github.com/flare-foundation/fsp-rewards.git ~/fsp-rewards
```
2. Set up .env file using env.example
3. Run the script
```
go run cmd/rewardmanager/main.go ../fsp-rewards/flare {start_epoch} {end_epoch}
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

## Roadmap
[ ] Add CSV report support