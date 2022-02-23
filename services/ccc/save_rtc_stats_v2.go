package ccc

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

// SaveRTCStatsV2 invokes the ccc.SaveRTCStatsV2 API synchronously
func (client *Client) SaveRTCStatsV2(request *SaveRTCStatsV2Request) (response *SaveRTCStatsV2Response, err error) {
	response = CreateSaveRTCStatsV2Response()
	err = client.DoAction(request, response)
	return
}

// SaveRTCStatsV2WithChan invokes the ccc.SaveRTCStatsV2 API asynchronously
func (client *Client) SaveRTCStatsV2WithChan(request *SaveRTCStatsV2Request) (<-chan *SaveRTCStatsV2Response, <-chan error) {
	responseChan := make(chan *SaveRTCStatsV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveRTCStatsV2(request)
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

// SaveRTCStatsV2WithCallback invokes the ccc.SaveRTCStatsV2 API asynchronously
func (client *Client) SaveRTCStatsV2WithCallback(request *SaveRTCStatsV2Request, callback func(response *SaveRTCStatsV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveRTCStatsV2Response
		var err error
		defer close(result)
		response, err = client.SaveRTCStatsV2(request)
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

// SaveRTCStatsV2Request is the request struct for api SaveRTCStatsV2
type SaveRTCStatsV2Request struct {
	*requests.RpcRequest
	CallId         string `position:"Query" name:"CallId"`
	SenderReport   string `position:"Query" name:"SenderReport"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	ReceiverReport string `position:"Query" name:"ReceiverReport"`
	GoogAddress    string `position:"Query" name:"GoogAddress"`
	GeneralInfo    string `position:"Query" name:"GeneralInfo"`
}

// SaveRTCStatsV2Response is the response struct for api SaveRTCStatsV2
type SaveRTCStatsV2Response struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RowCount       int64  `json:"RowCount" xml:"RowCount"`
	Success        bool   `json:"Success" xml:"Success"`
	TimeStamp      int64  `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateSaveRTCStatsV2Request creates a request to invoke SaveRTCStatsV2 API
func CreateSaveRTCStatsV2Request() (request *SaveRTCStatsV2Request) {
	request = &SaveRTCStatsV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "SaveRTCStatsV2", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveRTCStatsV2Response creates a response to parse from SaveRTCStatsV2 response
func CreateSaveRTCStatsV2Response() (response *SaveRTCStatsV2Response) {
	response = &SaveRTCStatsV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}