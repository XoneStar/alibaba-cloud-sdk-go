package vpc

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

// UpdateIpsecServer invokes the vpc.UpdateIpsecServer API synchronously
func (client *Client) UpdateIpsecServer(request *UpdateIpsecServerRequest) (response *UpdateIpsecServerResponse, err error) {
	response = CreateUpdateIpsecServerResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateIpsecServerWithChan invokes the vpc.UpdateIpsecServer API asynchronously
func (client *Client) UpdateIpsecServerWithChan(request *UpdateIpsecServerRequest) (<-chan *UpdateIpsecServerResponse, <-chan error) {
	responseChan := make(chan *UpdateIpsecServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateIpsecServer(request)
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

// UpdateIpsecServerWithCallback invokes the vpc.UpdateIpsecServer API asynchronously
func (client *Client) UpdateIpsecServerWithCallback(request *UpdateIpsecServerRequest, callback func(response *UpdateIpsecServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateIpsecServerResponse
		var err error
		defer close(result)
		response, err = client.UpdateIpsecServer(request)
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

// UpdateIpsecServerRequest is the request struct for api UpdateIpsecServer
type UpdateIpsecServerRequest struct {
	*requests.RpcRequest
	IkeConfig              string           `position:"Query" name:"IkeConfig"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	IpsecConfig            string           `position:"Query" name:"IpsecConfig"`
	Psk                    string           `position:"Query" name:"Psk"`
	LocalSubnet            string           `position:"Query" name:"LocalSubnet"`
	IDaaSInstanceId        string           `position:"Query" name:"IDaaSInstanceId"`
	EffectImmediately      requests.Boolean `position:"Query" name:"EffectImmediately"`
	ClientIpPool           string           `position:"Query" name:"ClientIpPool"`
	DryRun                 string           `position:"Query" name:"DryRun"`
	CallerBid              string           `position:"Query" name:"callerBid"`
	PskEnabled             requests.Boolean `position:"Query" name:"PskEnabled"`
	MultiFactorAuthEnabled requests.Boolean `position:"Query" name:"MultiFactorAuthEnabled"`
	IpsecServerName        string           `position:"Query" name:"IpsecServerName"`
	IpsecServerId          string           `position:"Query" name:"IpsecServerId"`
}

// UpdateIpsecServerResponse is the response struct for api UpdateIpsecServer
type UpdateIpsecServerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateIpsecServerRequest creates a request to invoke UpdateIpsecServer API
func CreateUpdateIpsecServerRequest() (request *UpdateIpsecServerRequest) {
	request = &UpdateIpsecServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UpdateIpsecServer", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateIpsecServerResponse creates a response to parse from UpdateIpsecServer response
func CreateUpdateIpsecServerResponse() (response *UpdateIpsecServerResponse) {
	response = &UpdateIpsecServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}