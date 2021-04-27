package edas

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

// SecretsItem is a nested struct in edas response
type SecretsItem struct {
	Name         string            `json:"Name" xml:"Name"`
	Namespace    string            `json:"Namespace" xml:"Namespace"`
	CreationTime string            `json:"CreationTime" xml:"CreationTime"`
	Type         string            `json:"Type" xml:"Type"`
	ClusterId    string            `json:"ClusterId" xml:"ClusterId"`
	ClusterName  string            `json:"ClusterName" xml:"ClusterName"`
	Data         []DataItem        `json:"Data" xml:"Data"`
	RelatedApps  []RelatedAppsItem `json:"RelatedApps" xml:"RelatedApps"`
}
