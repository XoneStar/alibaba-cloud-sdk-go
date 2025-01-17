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

// SendSmsVerifyCode invokes the dypnsapi.SendSmsVerifyCode API synchronously
func (client *Client) SendSmsVerifyCode(request *SendSmsVerifyCodeRequest) (response *SendSmsVerifyCodeResponse, err error) {
	response = CreateSendSmsVerifyCodeResponse()
	err = client.DoAction(request, response)
	return
}

// SendSmsVerifyCodeWithChan invokes the dypnsapi.SendSmsVerifyCode API asynchronously
func (client *Client) SendSmsVerifyCodeWithChan(request *SendSmsVerifyCodeRequest) (<-chan *SendSmsVerifyCodeResponse, <-chan error) {
	responseChan := make(chan *SendSmsVerifyCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendSmsVerifyCode(request)
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

// SendSmsVerifyCodeWithCallback invokes the dypnsapi.SendSmsVerifyCode API asynchronously
func (client *Client) SendSmsVerifyCodeWithCallback(request *SendSmsVerifyCodeRequest, callback func(response *SendSmsVerifyCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendSmsVerifyCodeResponse
		var err error
		defer close(result)
		response, err = client.SendSmsVerifyCode(request)
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

// SendSmsVerifyCodeRequest is the request struct for api SendSmsVerifyCode
type SendSmsVerifyCodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CountryCode          string           `position:"Query" name:"CountryCode"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	SmsUpExtendCode      string           `position:"Query" name:"SmsUpExtendCode"`
	SignName             string           `position:"Query" name:"SignName"`
	RouteName            string           `position:"Query" name:"RouteName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ValidTime            requests.Integer `position:"Query" name:"ValidTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ReturnVerifyCode     requests.Boolean `position:"Query" name:"ReturnVerifyCode"`
	CodeType             requests.Integer `position:"Query" name:"CodeType"`
	SchemeName           string           `position:"Query" name:"SchemeName"`
	DuplicatePolicy      requests.Integer `position:"Query" name:"DuplicatePolicy"`
	OutId                string           `position:"Query" name:"OutId"`
	Interval             requests.Integer `position:"Query" name:"Interval"`
	TemplateCode         string           `position:"Query" name:"TemplateCode"`
	TemplateParam        string           `position:"Query" name:"TemplateParam"`
	CodeLength           requests.Integer `position:"Query" name:"CodeLength"`
}

// SendSmsVerifyCodeResponse is the response struct for api SendSmsVerifyCode
type SendSmsVerifyCodeResponse struct {
	*responses.BaseResponse
	AccessDeniedDetail string `json:"AccessDeniedDetail" xml:"AccessDeniedDetail"`
	Message            string `json:"Message" xml:"Message"`
	Code               string `json:"Code" xml:"Code"`
	Success            bool   `json:"Success" xml:"Success"`
	Model              Model  `json:"Model" xml:"Model"`
}

// CreateSendSmsVerifyCodeRequest creates a request to invoke SendSmsVerifyCode API
func CreateSendSmsVerifyCodeRequest() (request *SendSmsVerifyCodeRequest) {
	request = &SendSmsVerifyCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "SendSmsVerifyCode", "", "")
	request.Method = requests.POST
	return
}

// CreateSendSmsVerifyCodeResponse creates a response to parse from SendSmsVerifyCode response
func CreateSendSmsVerifyCodeResponse() (response *SendSmsVerifyCodeResponse) {
	response = &SendSmsVerifyCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
