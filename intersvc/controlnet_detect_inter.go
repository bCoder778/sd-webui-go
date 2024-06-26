package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /controlnet/detect
type ControlnetDetect struct {
	RequestItem  *ControlnetDetectRequest
	ResponseItem *ControlnetDetectResponse
	Error        error
}

func (d *ControlnetDetect) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewDetectControlnetDetectPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.DetectControlnetDetectPost(RequestData)
	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(error(nil)) {
			d.Error = err
			return
		}
		errorValue := reflect.ValueOf(err).Elem().FieldByName("Payload")
		if !errorValue.IsValid() {
			d.Error = err
			return
		}
		d.Error = fmt.Errorf("%v", errorValue.Elem())
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &ControlnetDetectResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*ControlnetDetectResponse)
}

func (d *ControlnetDetect) GetResponse() *ControlnetDetectResponse {
	return d.ResponseItem
}
