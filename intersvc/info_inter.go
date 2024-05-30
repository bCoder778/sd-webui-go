package intersvc

import (
	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /info
type Info struct {
	ResponseItem *InfoResponse
	Error        error
}

func (d *Info) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewAPIInfoInfoGetParams()
	ResponseData, err := inter.Client.Operations.APIInfoInfoGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfoResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfoResponse)
}

func (d *Info) GetResponse() *InfoResponse {
	return d.ResponseItem
}
