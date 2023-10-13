package push

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

// DeviceInfo is a nested struct in push response
type DeviceInfo struct {
	Account        string `json:"Account" xml:"Account"`
	LastOnlineTime string `json:"LastOnlineTime" xml:"LastOnlineTime"`
	PhoneNumber    string `json:"PhoneNumber" xml:"PhoneNumber"`
	PushEnabled    bool   `json:"PushEnabled" xml:"PushEnabled"`
	DeviceType     string `json:"DeviceType" xml:"DeviceType"`
	DeviceId       string `json:"DeviceId" xml:"DeviceId"`
	Online         bool   `json:"Online" xml:"Online"`
	Tags           string `json:"Tags" xml:"Tags"`
	DeviceToken    string `json:"DeviceToken" xml:"DeviceToken"`
	Alias          string `json:"Alias" xml:"Alias"`
	Brand          string `json:"Brand" xml:"Brand"`
	Model          string `json:"Model" xml:"Model"`
}
