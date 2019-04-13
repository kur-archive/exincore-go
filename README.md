# exincore-go

ExinCore golang client

## Installation

```bash
go get https://github.com/Kurisu-package/exincore-go
```

## Example

**Create Order**

`usdt` => `eth`

```go
var userId, sessionId, pin, pinToken, priKey string
usdt := "815b0b1a-2764-3736-8faa-42d694fa620a"
eth := "43d61dcd-e413-450d-80b8-101d5e903357"

client := NewExinCoreClient(userId, sessionId, pin, pinToken, priKey)

client.CreateOrder(context.TODO(), usdt, eth, "", 1)
```

**ReadPair**

```go
// Just Read base coin is xin pairs

base := "c94ac88f-4671-3976-b60a-09064f1811e8"

client := NewExinCoreClient("", "", "", "", "")
info, err := client.ReadPair(base, "")

// Reading base coin is xin and exchange coin is btc pair
base := "c94ac88f-4671-3976-b60a-09064f1811e8"
exchange := "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
client := NewExinCoreClient("", "", "", "", "")
info, err := client.ReadPair(base, "")
```