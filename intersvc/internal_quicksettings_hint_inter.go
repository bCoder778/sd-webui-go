package intersvc

import (
	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /internal/quicksettings-hint
type InternalQuicksettingsHint struct {
	ResponseItem *InternalQuicksettingsHintResponse
	Error        error
}

func (d *InternalQuicksettingsHint) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewQuicksettingsHintInternalQuicksettingsHintGetParams()
	ResponseData, err := inter.Client.Operations.QuicksettingsHintInternalQuicksettingsHintGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InternalQuicksettingsHintResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InternalQuicksettingsHintResponse)
}

func (d *InternalQuicksettingsHint) GetResponse() *InternalQuicksettingsHintResponse {
	return d.ResponseItem
}
