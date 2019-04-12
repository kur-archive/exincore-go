package exincore_go

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	bot "github.com/MixinNetwork/bot-api-go-client"
	"github.com/MixinNetwork/go-number"
	"github.com/gofrs/uuid"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"net/http"
)

const (
	exincoreUri = "https://exinone.com/exincore/"
	exincore    = "61103d28-3ac2-44a2-ae34-bd956070dab1"
)

type EClient struct {
	userUuid    string
	sessionUuid string
	pin         string
	pinToken    string
	privateKey  string
}

func New(userUuid, sessionUuid, pin, pinToken, privateKey string) *EClient {
	return &EClient{
		userUuid:    userUuid,
		sessionUuid: sessionUuid,
		pin:         pin,
		pinToken:    pinToken,
		privateKey:  privateKey,
	}
}

type PairInfo struct {
	BaseAsset           string   `json:"base_asset"`
	BaseAssetSymbol     string   `json:"base_asset_symbol"`
	ExchangeAsset       string   `json:"exchange_asset"`
	ExchangeAssetSymbol string   `json:"exchange_asset_symbol"`
	MinimumAmount       string   `json:"minimum_amount"`
	MaximumAmount       string   `json:"maximum_amount"`
	Exchanges           []string `json:"exchanges"`
	Price               string   `json:"price"`
}

// ----------------------

func (e *EClient) CreateOrder(ctx context.Context, base, exchange, traceId string, amount float64) error {
	packUuid, _ := uuid.FromString(exchange)
	var OrderAction struct {
		A uuid.UUID
	}
	OrderAction.A = packUuid
	pack, _ := msgpack.Marshal(OrderAction)
	memo := base64.StdEncoding.EncodeToString(pack)

	err := bot.CreateTransfer(ctx, &bot.TransferInput{
		AssetId:     base,
		RecipientId: exincore,
		Amount:      number.FromFloat(amount),
		TraceId:     traceId,
		Memo:        memo,
	}, e.userUuid, e.sessionUuid, e.privateKey, e.pin, e.pinToken)

	if err != nil {
		return err
	}

	return nil
}

func (e *EClient) ReadPair(base, exchange string) (*[]PairInfo, error) {
	url := exincoreUri + "markets?base_asset=" + base + "&exchange_asset=" + exchange
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Code    int        `json:"code"`
		Data    []PairInfo `json:"data"`
		Message string     `json:"message"`
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 || resp.Message != "success" {
		return nil, errors.New(resp.Message)
	}

	return &resp.Data, nil
}
