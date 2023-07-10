package rds

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

// ModifyDBInstanceSpec invokes the rds.ModifyDBInstanceSpec API synchronously
func (client *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
	response = CreateModifyDBInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceSpecWithChan invokes the rds.ModifyDBInstanceSpec API asynchronously
func (client *Client) ModifyDBInstanceSpecWithChan(request *ModifyDBInstanceSpecRequest) (<-chan *ModifyDBInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceSpec(request)
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

// ModifyDBInstanceSpecWithCallback invokes the rds.ModifyDBInstanceSpec API asynchronously
func (client *Client) ModifyDBInstanceSpecWithCallback(request *ModifyDBInstanceSpecRequest, callback func(response *ModifyDBInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceSpec(request)
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

// ModifyDBInstanceSpecRequest is the request struct for api ModifyDBInstanceSpec
type ModifyDBInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer                            `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage       requests.Integer                            `position:"Query" name:"DBInstanceStorage"`
	ClientToken             string                                      `position:"Query" name:"ClientToken"`
	EngineVersion           string                                      `position:"Query" name:"EngineVersion"`
	AutoUseCoupon           requests.Boolean                            `position:"Query" name:"AutoUseCoupon"`
	ResourceGroupId         string                                      `position:"Query" name:"ResourceGroupId"`
	ServerlessConfiguration ModifyDBInstanceSpecServerlessConfiguration `position:"Query" name:"ServerlessConfiguration"  type:"Struct"`
	EffectiveTime           string                                      `position:"Query" name:"EffectiveTime"`
	DBInstanceTransType     string                                      `position:"Query" name:"DBInstanceTransType"`
	DBInstanceId            string                                      `position:"Query" name:"DBInstanceId"`
	SwitchTime              string                                      `position:"Query" name:"SwitchTime"`
	DBInstanceStorageType   string                                      `position:"Query" name:"DBInstanceStorageType"`
	SourceBiz               string                                      `position:"Query" name:"SourceBiz"`
	DedicatedHostGroupId    string                                      `position:"Query" name:"DedicatedHostGroupId"`
	Direction               string                                      `position:"Query" name:"Direction"`
	ResourceOwnerAccount    string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string                                      `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer                            `position:"Query" name:"OwnerId"`
	UsedTime                requests.Integer                            `position:"Query" name:"UsedTime"`
	BurstingEnabled         requests.Boolean                            `position:"Query" name:"BurstingEnabled"`
	DBInstanceClass         string                                      `position:"Query" name:"DBInstanceClass"`
	ZoneId                  string                                      `position:"Query" name:"ZoneId"`
	Category                string                                      `position:"Query" name:"Category"`
	PayType                 string                                      `position:"Query" name:"PayType"`
}

// ModifyDBInstanceSpecServerlessConfiguration is a repeated param struct in ModifyDBInstanceSpecRequest
type ModifyDBInstanceSpecServerlessConfiguration struct {
	MinCapacity string `name:"MinCapacity"`
	MaxCapacity string `name:"MaxCapacity"`
	AutoPause   string `name:"AutoPause"`
	SwitchForce string `name:"SwitchForce"`
}

// ModifyDBInstanceSpecResponse is the response struct for api ModifyDBInstanceSpec
type ModifyDBInstanceSpecResponse struct {
	*responses.BaseResponse
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	OrderId      int64  `json:"OrderId" xml:"OrderId"`
}

// CreateModifyDBInstanceSpecRequest creates a request to invoke ModifyDBInstanceSpec API
func CreateModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
	request = &ModifyDBInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSpec", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceSpecResponse creates a response to parse from ModifyDBInstanceSpec response
func CreateModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
	response = &ModifyDBInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
