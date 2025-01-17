package dypnsapi

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

// CreateVerifyScheme invokes the dypnsapi.CreateVerifyScheme API synchronously
func (client *Client) CreateVerifyScheme(request *CreateVerifySchemeRequest) (response *CreateVerifySchemeResponse, err error) {
	response = CreateCreateVerifySchemeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVerifySchemeWithChan invokes the dypnsapi.CreateVerifyScheme API asynchronously
func (client *Client) CreateVerifySchemeWithChan(request *CreateVerifySchemeRequest) (<-chan *CreateVerifySchemeResponse, <-chan error) {
	responseChan := make(chan *CreateVerifySchemeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVerifyScheme(request)
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

// CreateVerifySchemeWithCallback invokes the dypnsapi.CreateVerifyScheme API asynchronously
func (client *Client) CreateVerifySchemeWithCallback(request *CreateVerifySchemeRequest, callback func(response *CreateVerifySchemeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVerifySchemeResponse
		var err error
		defer close(result)
		response, err = client.CreateVerifyScheme(request)
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

// CreateVerifySchemeRequest is the request struct for api CreateVerifyScheme
type CreateVerifySchemeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Origin               string           `position:"Query" name:"Origin"`
	BundleId             string           `position:"Query" name:"BundleId"`
	AuthType             string           `position:"Query" name:"AuthType"`
	AppName              string           `position:"Query" name:"AppName"`
	IpWhiteList          string           `position:"Query" name:"IpWhiteList"`
	RouteName            string           `position:"Query" name:"RouteName"`
	Email                string           `position:"Query" name:"Email"`
	PackSign             string           `position:"Query" name:"PackSign"`
	PackName             string           `position:"Query" name:"PackName"`
	CuApiCode            requests.Integer `position:"Query" name:"CuApiCode"`
	SceneType            string           `position:"Query" name:"SceneType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CtApiCode            requests.Integer `position:"Query" name:"CtApiCode"`
	OsType               string           `position:"Query" name:"OsType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Url                  string           `position:"Query" name:"Url"`
	CmApiCode            requests.Integer `position:"Query" name:"CmApiCode"`
	SchemeName           string           `position:"Query" name:"SchemeName"`
	SmsSignName          string           `position:"Query" name:"SmsSignName"`
}

// CreateVerifySchemeResponse is the response struct for api CreateVerifyScheme
type CreateVerifySchemeResponse struct {
	*responses.BaseResponse
	Message             string              `json:"Message" xml:"Message"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	HttpStatusCode      int64               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                string              `json:"Code" xml:"Code"`
	Success             bool                `json:"Success" xml:"Success"`
	GateVerifySchemeDTO GateVerifySchemeDTO `json:"GateVerifySchemeDTO" xml:"GateVerifySchemeDTO"`
}

// CreateCreateVerifySchemeRequest creates a request to invoke CreateVerifyScheme API
func CreateCreateVerifySchemeRequest() (request *CreateVerifySchemeRequest) {
	request = &CreateVerifySchemeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "CreateVerifyScheme", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateVerifySchemeResponse creates a response to parse from CreateVerifyScheme response
func CreateCreateVerifySchemeResponse() (response *CreateVerifySchemeResponse) {
	response = &CreateVerifySchemeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
