package objectdet

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

// DetectVehicleICongestion invokes the objectdet.DetectVehicleICongestion API synchronously
func (client *Client) DetectVehicleICongestion(request *DetectVehicleICongestionRequest) (response *DetectVehicleICongestionResponse, err error) {
	response = CreateDetectVehicleICongestionResponse()
	err = client.DoAction(request, response)
	return
}

// DetectVehicleICongestionWithChan invokes the objectdet.DetectVehicleICongestion API asynchronously
func (client *Client) DetectVehicleICongestionWithChan(request *DetectVehicleICongestionRequest) (<-chan *DetectVehicleICongestionResponse, <-chan error) {
	responseChan := make(chan *DetectVehicleICongestionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectVehicleICongestion(request)
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

// DetectVehicleICongestionWithCallback invokes the objectdet.DetectVehicleICongestion API asynchronously
func (client *Client) DetectVehicleICongestionWithCallback(request *DetectVehicleICongestionRequest, callback func(response *DetectVehicleICongestionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectVehicleICongestionResponse
		var err error
		defer close(result)
		response, err = client.DetectVehicleICongestion(request)
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

// DetectVehicleICongestionRequest is the request struct for api DetectVehicleICongestion
type DetectVehicleICongestionRequest struct {
	*requests.RpcRequest
	RoadRegions                string `position:"Body" name:"RoadRegions"`
	PreRegionIntersectFeatures string `position:"Body" name:"PreRegionIntersectFeatures"`
	ImageURL                   string `position:"Body" name:"ImageURL"`
}

// DetectVehicleICongestionResponse is the response struct for api DetectVehicleICongestion
type DetectVehicleICongestionResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectVehicleICongestionRequest creates a request to invoke DetectVehicleICongestion API
func CreateDetectVehicleICongestionRequest() (request *DetectVehicleICongestionRequest) {
	request = &DetectVehicleICongestionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("objectdet", "2019-12-30", "DetectVehicleICongestion", "objectdet", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectVehicleICongestionResponse creates a response to parse from DetectVehicleICongestion response
func CreateDetectVehicleICongestionResponse() (response *DetectVehicleICongestionResponse) {
	response = &DetectVehicleICongestionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
