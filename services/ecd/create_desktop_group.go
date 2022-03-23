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

// CreateDesktopGroup invokes the ecd.CreateDesktopGroup API synchronously
func (client *Client) CreateDesktopGroup(request *CreateDesktopGroupRequest) (response *CreateDesktopGroupResponse, err error) {
	response = CreateCreateDesktopGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDesktopGroupWithChan invokes the ecd.CreateDesktopGroup API asynchronously
func (client *Client) CreateDesktopGroupWithChan(request *CreateDesktopGroupRequest) (<-chan *CreateDesktopGroupResponse, <-chan error) {
	responseChan := make(chan *CreateDesktopGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDesktopGroup(request)
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

// CreateDesktopGroupWithCallback invokes the ecd.CreateDesktopGroup API asynchronously
func (client *Client) CreateDesktopGroupWithCallback(request *CreateDesktopGroupRequest, callback func(response *CreateDesktopGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDesktopGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateDesktopGroup(request)
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

// CreateDesktopGroupRequest is the request struct for api CreateDesktopGroup
type CreateDesktopGroupRequest struct {
	*requests.RpcRequest
	OfficeSiteId            string           `position:"Query" name:"OfficeSiteId"`
	EndUserIds              *[]string        `position:"Query" name:"EndUserIds"  type:"Repeated"`
	ScaleStrategyId         string           `position:"Query" name:"ScaleStrategyId"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	BundleId                string           `position:"Query" name:"BundleId"`
	BindAmount              requests.Integer `position:"Query" name:"BindAmount"`
	LoadPolicy              requests.Integer `position:"Query" name:"LoadPolicy"`
	DesktopGroupName        string           `position:"Query" name:"DesktopGroupName"`
	AllowBufferCount        requests.Integer `position:"Query" name:"AllowBufferCount"`
	DefaultInitDesktopCount requests.Integer `position:"Query" name:"DefaultInitDesktopCount"`
	DirectoryId             string           `position:"Query" name:"DirectoryId"`
	MinDesktopsCount        requests.Integer `position:"Query" name:"MinDesktopsCount"`
	MaxDesktopsCount        requests.Integer `position:"Query" name:"MaxDesktopsCount"`
	Period                  requests.Integer `position:"Query" name:"Period"`
	AllowAutoSetup          requests.Integer `position:"Query" name:"AllowAutoSetup"`
	AutoPay                 requests.Boolean `position:"Query" name:"AutoPay"`
	Comments                string           `position:"Query" name:"Comments"`
	ResetType               requests.Integer `position:"Query" name:"ResetType"`
	OwnType                 requests.Integer `position:"Query" name:"OwnType"`
	KeepDuration            requests.Integer `position:"Query" name:"KeepDuration"`
	PeriodUnit              string           `position:"Query" name:"PeriodUnit"`
	VpcId                   string           `position:"Query" name:"VpcId"`
	ChargeType              string           `position:"Query" name:"ChargeType"`
	PolicyGroupId           string           `position:"Query" name:"PolicyGroupId"`
}

// CreateDesktopGroupResponse is the response struct for api CreateDesktopGroup
type CreateDesktopGroupResponse struct {
	*responses.BaseResponse
	DesktopGroupId string   `json:"DesktopGroupId" xml:"DesktopGroupId"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	OrderIds       []string `json:"OrderIds" xml:"OrderIds"`
}

// CreateCreateDesktopGroupRequest creates a request to invoke CreateDesktopGroup API
func CreateCreateDesktopGroupRequest() (request *CreateDesktopGroupRequest) {
	request = &CreateDesktopGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateDesktopGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDesktopGroupResponse creates a response to parse from CreateDesktopGroup response
func CreateCreateDesktopGroupResponse() (response *CreateDesktopGroupResponse) {
	response = &CreateDesktopGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}