package intersvc

import (
	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /token
type Token struct {
	ResponseItem *TokenResponse
	Error        error
}

func (d *Token) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetTokenTokenGetParams()
	ResponseData, err := inter.Client.Operations.GetTokenTokenGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &TokenResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*TokenResponse)
}

func (d *Token) GetResponse() *TokenResponse {
	return d.ResponseItem
}
