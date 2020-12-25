package imageprocess

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

// ScreenChestCT invokes the imageprocess.ScreenChestCT API synchronously
func (client *Client) ScreenChestCT(request *ScreenChestCTRequest) (response *ScreenChestCTResponse, err error) {
	response = CreateScreenChestCTResponse()
	err = client.DoAction(request, response)
	return
}

// ScreenChestCTWithChan invokes the imageprocess.ScreenChestCT API asynchronously
func (client *Client) ScreenChestCTWithChan(request *ScreenChestCTRequest) (<-chan *ScreenChestCTResponse, <-chan error) {
	responseChan := make(chan *ScreenChestCTResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScreenChestCT(request)
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

// ScreenChestCTWithCallback invokes the imageprocess.ScreenChestCT API asynchronously
func (client *Client) ScreenChestCTWithCallback(request *ScreenChestCTRequest, callback func(response *ScreenChestCTResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScreenChestCTResponse
		var err error
		defer close(result)
		response, err = client.ScreenChestCT(request)
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

// ScreenChestCTRequest is the request struct for api ScreenChestCT
type ScreenChestCTRequest struct {
	*requests.RpcRequest
	DataFormat string                  `position:"Body" name:"DataFormat"`
	URLList    *[]ScreenChestCTURLList `position:"Body" name:"URLList"  type:"Repeated"`
	OrgId      string                  `position:"Body" name:"OrgId"`
	Async      requests.Boolean        `position:"Body" name:"Async"`
	OrgName    string                  `position:"Body" name:"OrgName"`
}

// ScreenChestCTURLList is a repeated param struct in ScreenChestCTRequest
type ScreenChestCTURLList struct {
	URL string `name:"URL"`
}

// ScreenChestCTResponse is the response struct for api ScreenChestCT
type ScreenChestCTResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateScreenChestCTRequest creates a request to invoke ScreenChestCT API
func CreateScreenChestCTRequest() (request *ScreenChestCTRequest) {
	request = &ScreenChestCTRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "ScreenChestCT", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateScreenChestCTResponse creates a response to parse from ScreenChestCT response
func CreateScreenChestCTResponse() (response *ScreenChestCTResponse) {
	response = &ScreenChestCTResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}