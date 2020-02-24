package outboundbot

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

// ModifyScript invokes the outboundbot.ModifyScript API synchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyscript.html
func (client *Client) ModifyScript(request *ModifyScriptRequest) (response *ModifyScriptResponse, err error) {
	response = CreateModifyScriptResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScriptWithChan invokes the outboundbot.ModifyScript API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScriptWithChan(request *ModifyScriptRequest) (<-chan *ModifyScriptResponse, <-chan error) {
	responseChan := make(chan *ModifyScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScript(request)
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

// ModifyScriptWithCallback invokes the outboundbot.ModifyScript API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/modifyscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScriptWithCallback(request *ModifyScriptRequest, callback func(response *ModifyScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScriptResponse
		var err error
		defer close(result)
		response, err = client.ModifyScript(request)
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

// ModifyScriptRequest is the request struct for api ModifyScript
type ModifyScriptRequest struct {
	*requests.RpcRequest
	Industry          string `position:"Query" name:"Industry"`
	ScriptName        string `position:"Query" name:"ScriptName"`
	Scene             string `position:"Query" name:"Scene"`
	ScriptId          string `position:"Query" name:"ScriptId"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	ScriptDescription string `position:"Query" name:"ScriptDescription"`
}

// ModifyScriptResponse is the response struct for api ModifyScript
type ModifyScriptResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Script         Script `json:"Script" xml:"Script"`
}

// CreateModifyScriptRequest creates a request to invoke ModifyScript API
func CreateModifyScriptRequest() (request *ModifyScriptRequest) {
	request = &ModifyScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyScript", "outboundbot", "openAPI")
	return
}

// CreateModifyScriptResponse creates a response to parse from ModifyScript response
func CreateModifyScriptResponse() (response *ModifyScriptResponse) {
	response = &ModifyScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}