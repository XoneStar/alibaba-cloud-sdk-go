package live

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

// GetMultiRateConfig invokes the live.GetMultiRateConfig API synchronously
func (client *Client) GetMultiRateConfig(request *GetMultiRateConfigRequest) (response *GetMultiRateConfigResponse, err error) {
	response = CreateGetMultiRateConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetMultiRateConfigWithChan invokes the live.GetMultiRateConfig API asynchronously
func (client *Client) GetMultiRateConfigWithChan(request *GetMultiRateConfigRequest) (<-chan *GetMultiRateConfigResponse, <-chan error) {
	responseChan := make(chan *GetMultiRateConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMultiRateConfig(request)
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

// GetMultiRateConfigWithCallback invokes the live.GetMultiRateConfig API asynchronously
func (client *Client) GetMultiRateConfigWithCallback(request *GetMultiRateConfigRequest, callback func(response *GetMultiRateConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMultiRateConfigResponse
		var err error
		defer close(result)
		response, err = client.GetMultiRateConfig(request)
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

// GetMultiRateConfigRequest is the request struct for api GetMultiRateConfig
type GetMultiRateConfigRequest struct {
	*requests.RpcRequest
	App        string           `position:"Query" name:"App"`
	GroupId    string           `position:"Query" name:"GroupId"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// GetMultiRateConfigResponse is the response struct for api GetMultiRateConfig
type GetMultiRateConfigResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Message       string        `json:"Message" xml:"Message"`
	Code          int           `json:"Code" xml:"Code"`
	Domain        string        `json:"Domain" xml:"Domain"`
	App           string        `json:"App" xml:"App"`
	Stream        string        `json:"Stream" xml:"Stream"`
	AvFormat      string        `json:"AvFormat" xml:"AvFormat"`
	GroupId       string        `json:"GroupId" xml:"GroupId"`
	IsLazy        string        `json:"IsLazy" xml:"IsLazy"`
	IsTimeAlign   string        `json:"IsTimeAlign" xml:"IsTimeAlign"`
	TemplatesInfo TemplatesInfo `json:"TemplatesInfo" xml:"TemplatesInfo"`
}

// CreateGetMultiRateConfigRequest creates a request to invoke GetMultiRateConfig API
func CreateGetMultiRateConfigRequest() (request *GetMultiRateConfigRequest) {
	request = &GetMultiRateConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "GetMultiRateConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMultiRateConfigResponse creates a response to parse from GetMultiRateConfig response
func CreateGetMultiRateConfigResponse() (response *GetMultiRateConfigResponse) {
	response = &GetMultiRateConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
