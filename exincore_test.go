package exincore_go

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type keys struct {
	MixinID      string `json:"MixinId"`
	ClientID     string `json:"ClientId"`
	ClientSecret string `json:"ClientSecret"`
	Pin          string `json:"Pin"`
	PinToken     string `json:"PinToken"`
	SessionID    string `json:"SessionId"`
	PrivateKey   string `json:"PrivateKey"`
}

func TestCreateOrder(t *testing.T) {
	file, err := os.Open("key.txt")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var keys keys
	err = json.Unmarshal(data, &keys)
	if err != nil {
		panic(err)
	}

	client := NewExinCoreClient(keys.ClientID, keys.SessionID, keys.Pin, keys.PinToken, keys.PrivateKey)
	err = client.CreateOrder(context.TODO(), "815b0b1a-2764-3736-8faa-42d694fa620a", "43d61dcd-e413-450d-80b8-101d5e903357", "", 1)
	if err != nil {
		panic(err)
	}

}

func TestReadPair(t *testing.T) {
	base := "c94ac88f-4671-3976-b60a-09064f1811e8"
	exchange := "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
	client := NewExinCoreClient("", "", "", "", "")
	info, err := client.ReadPair(base, "")
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(*info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bytes)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	base = "c94ac88f-4671-3976-b60a-09064f1811e8"
	exchange = "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
	info, err = client.ReadPair(base, exchange)
	if err != nil {
		panic(err)
	}
	bytes, err = json.Marshal(*info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bytes)
}
