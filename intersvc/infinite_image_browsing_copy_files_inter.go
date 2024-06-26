package intersvc

import (
	"fmt"
	"reflect"

	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/copy_files
type InfiniteImageBrowsingCopyFiles struct {
	RequestItem  *InfiniteImageBrowsingCopyFilesRequest
	ResponseItem *InfiniteImageBrowsingCopyFilesResponse
	Error        error
}

func (d *InfiniteImageBrowsingCopyFiles) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewCopyFilesInfiniteImageBrowsingCopyFilesPostParams()
	RequestData.Body = d.RequestItem
	ResponseData, err := inter.Client.Operations.CopyFilesInfiniteImageBrowsingCopyFilesPost(RequestData)
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
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingCopyFilesResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingCopyFilesResponse)
}

func (d *InfiniteImageBrowsingCopyFiles) GetResponse() *InfiniteImageBrowsingCopyFilesResponse {
	return d.ResponseItem
}
