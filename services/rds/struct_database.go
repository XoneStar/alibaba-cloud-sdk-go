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

// Database is a nested struct in rds response
type Database struct {
	CharacterSetName string                      `json:"CharacterSetName" xml:"CharacterSetName"`
	Engine           string                      `json:"Engine" xml:"Engine"`
	DBStatus         string                      `json:"DBStatus" xml:"DBStatus"`
	TDEStatus        string                      `json:"TDEStatus" xml:"TDEStatus"`
	ConnLimit        string                      `json:"ConnLimit" xml:"ConnLimit"`
	PageSize         int                         `json:"PageSize" xml:"PageSize"`
	ResourceGroupId  string                      `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tablespace       string                      `json:"Tablespace" xml:"Tablespace"`
	TotalCount       int                         `json:"TotalCount" xml:"TotalCount"`
	DBName           string                      `json:"DBName" xml:"DBName"`
	Ctype            string                      `json:"Ctype" xml:"Ctype"`
	PageNumber       int                         `json:"PageNumber" xml:"PageNumber"`
	DBInstanceId     string                      `json:"DBInstanceId" xml:"DBInstanceId"`
	DBDescription    string                      `json:"DBDescription" xml:"DBDescription"`
	Collate          string                      `json:"Collate" xml:"Collate"`
	BasicInfo        BasicInfo                   `json:"BasicInfo" xml:"BasicInfo"`
	RuntimeInfo      RuntimeInfo                 `json:"RuntimeInfo" xml:"RuntimeInfo"`
	AdvancedInfo     AdvancedInfo                `json:"AdvancedInfo" xml:"AdvancedInfo"`
	Accounts         AccountsInDescribeDatabases `json:"Accounts" xml:"Accounts"`
}
