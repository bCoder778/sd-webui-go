package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/check_path_exists
type InfiniteImageBrowsingCheckPathExists struct {
	RequestItem  *InfiniteImageBrowsingCheckPathExistsRequest
	ResponseItem *InfiniteImageBrowsingCheckPathExistsResponse
	Error        error
}

func (d *InfiniteImageBrowsingCheckPathExists) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCheckPathExistsInfiniteImageBrowsingCheckPathExistsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CheckPathExistsInfiniteImageBrowsingCheckPathExistsPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingCheckPathExistsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingCheckPathExistsResponse)
}

func (d *InfiniteImageBrowsingCheckPathExists) GetResponse() *InfiniteImageBrowsingCheckPathExistsResponse {
	return d.ResponseItem
}
