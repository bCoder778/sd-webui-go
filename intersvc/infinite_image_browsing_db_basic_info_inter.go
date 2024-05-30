package intersvc

import (
	webui "github.com/bCoder778/sd-webui-go"
	SdApiOperation "github.com/bCoder778/sd-webui-go/stablediffusion/client/operations"
)

// API Path: /infinite_image_browsing/db/basic_info
type InfiniteImageBrowsingDbBasicInfo struct {
	ResponseItem *InfiniteImageBrowsingDbBasicInfoResponse
	Error        error
}

func (d *InfiniteImageBrowsingDbBasicInfo) Action(inter *webui.StableDiffInterface) {
	RequestData := SdApiOperation.NewGetDbBasicInfoInfiniteImageBrowsingDbBasicInfoGetParams()
	ResponseData, err := inter.Client.Operations.GetDbBasicInfoInfiniteImageBrowsingDbBasicInfoGet(RequestData)
	if err != nil {
		d.Error = err
		return
	}
	ResponseItem, err := ConvertResponse(ResponseData.Payload, &InfiniteImageBrowsingDbBasicInfoResponse{})
	if err != nil {
		d.Error = err
		return
	}
	d.ResponseItem = ResponseItem.(*InfiniteImageBrowsingDbBasicInfoResponse)
}

func (d *InfiniteImageBrowsingDbBasicInfo) GetResponse() *InfiniteImageBrowsingDbBasicInfoResponse {
	return d.ResponseItem
}
