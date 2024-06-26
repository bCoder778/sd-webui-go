package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/scanned_paths
type InfiniteImageBrowsingDbScannedPaths struct {
	RequestItem  *InfiniteImageBrowsingDbScannedPathsRequest
	ResponseItem *InfiniteImageBrowsingDbScannedPathsResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbScannedPaths) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCreateScannedPathInfiniteImageBrowsingDbScannedPathsPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CreateScannedPathInfiniteImageBrowsingDbScannedPathsPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbScannedPathsResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbScannedPathsResponse)
}

func (d *InfiniteImageBrowsingDbScannedPaths) GetResponse() *InfiniteImageBrowsingDbScannedPathsResponse {
	return d.ResponseItem
}
