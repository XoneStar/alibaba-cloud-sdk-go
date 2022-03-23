package ecd

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDesktopSpec invokes the ecd.ModifyDesktopSpec API synchronously
func (client *Client) ModifyDesktopSpec(request *ModifyDesktopSpecRequest) (response *ModifyDesktopSpecResponse, err error) {
	response = CreateModifyDesktopSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDesktopSpecWithChan invokes the ecd.ModifyDesktopSpec API asynchronously
func (client *Client) ModifyDesktopSpecWithChan(request *ModifyDesktopSpecRequest) (<-chan *ModifyDesktopSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyDesktopSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDesktopSpec(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDesktopSpecWithCallback invokes the ecd.ModifyDesktopSpec API asynchronously
func (client *Client) ModifyDesktopSpecWithCallback(request *ModifyDesktopSpecRequest, callback func(response *ModifyDesktopSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDesktopSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyDesktopSpec(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDesktopSpecRequest is the request struct for api ModifyDesktopSpec
type ModifyDesktopSpecRequest struct {
	*requests.RpcRequest
	AutoPay                  requests.Boolean `position:"Query" name:"AutoPay"`
	UserDiskPerformanceLevel string           `position:"Query" name:"UserDiskPerformanceLevel"`
	PromotionId              string           `position:"Query" name:"PromotionId"`
	UserDiskSizeGib          requests.Integer `position:"Query" name:"UserDiskSizeGib"`
	DesktopId                string           `position:"Query" name:"DesktopId"`
	DesktopType              string           `position:"Query" name:"DesktopType"`
	RootDiskSizeGib          requests.Integer `position:"Query" name:"RootDiskSizeGib"`
}

// ModifyDesktopSpecResponse is the response struct for api ModifyDesktopSpec
type ModifyDesktopSpecResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDesktopSpecRequest creates a request to invoke ModifyDesktopSpec API
func CreateModifyDesktopSpecRequest() (request *ModifyDesktopSpecRequest) {
	request = &ModifyDesktopSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyDesktopSpec", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDesktopSpecResponse creates a response to parse from ModifyDesktopSpec response
func CreateModifyDesktopSpecResponse() (response *ModifyDesktopSpecResponse) {
	response = &ModifyDesktopSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}