package r_kvstore

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

// ItemsItem is a nested struct in r_kvstore response
type ItemsItem struct {
	ResultInfo            string         `json:"ResultInfo" xml:"ResultInfo"`
	InsName               string         `json:"InsName" xml:"InsName"`
	ConfictReason         string         `json:"ConfictReason" xml:"ConfictReason"`
	SourceDetail          string         `json:"SourceDetail" xml:"SourceDetail"`
	InternetIP            string         `json:"InternetIP" xml:"InternetIP"`
	TaskParams            string         `json:"TaskParams" xml:"TaskParams"`
	InconsistentFields    string         `json:"InconsistentFields" xml:"InconsistentFields"`
	InstanceIdA           string         `json:"InstanceIdA" xml:"InstanceIdA"`
	Deadline              string         `json:"Deadline" xml:"Deadline"`
	ConflictGtid          string         `json:"ConflictGtid" xml:"ConflictGtid"`
	DestinationDetail     string         `json:"DestinationDetail" xml:"DestinationDetail"`
	DbType                string         `json:"DbType" xml:"DbType"`
	Key                   string         `json:"Key" xml:"Key"`
	ModifiedTime          string         `json:"ModifiedTime" xml:"ModifiedTime"`
	CreatedTime           string         `json:"CreatedTime" xml:"CreatedTime"`
	RegionId              string         `json:"RegionId" xml:"RegionId"`
	InstanceNetworkType   string         `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	RecoveryMode          string         `json:"RecoveryMode" xml:"RecoveryMode"`
	AbnormalType          string         `json:"AbnormalType" xml:"AbnormalType"`
	SwitchTime            string         `json:"SwitchTime" xml:"SwitchTime"`
	DBInstanceId          string         `json:"DBInstanceId" xml:"DBInstanceId"`
	DatabaseName          string         `json:"DatabaseName" xml:"DatabaseName"`
	DetailInfo            string         `json:"DetailInfo" xml:"DetailInfo"`
	SecurityIPList        string         `json:"SecurityIPList" xml:"SecurityIPList"`
	Id                    int            `json:"Id" xml:"Id"`
	Status                int            `json:"Status" xml:"Status"`
	Schema                string         `json:"Schema" xml:"Schema"`
	ReadWriteType         string         `json:"ReadWriteType" xml:"ReadWriteType"`
	DestinationInstanceId string         `json:"DestinationInstanceId" xml:"DestinationInstanceId"`
	HasInternetIP         bool           `json:"HasInternetIP" xml:"HasInternetIP"`
	PrepareInterval       string         `json:"PrepareInterval" xml:"PrepareInterval"`
	OccurTime             string         `json:"OccurTime" xml:"OccurTime"`
	InstanceIdB           string         `json:"InstanceIdB" xml:"InstanceIdB"`
	StartTime             string         `json:"StartTime" xml:"StartTime"`
	SourceInstanceId      string         `json:"SourceInstanceId" xml:"SourceInstanceId"`
	InconsistentType      string         `json:"InconsistentType" xml:"InconsistentType"`
	ConfictKey            string         `json:"ConfictKey" xml:"ConfictKey"`
	KeyType               string         `json:"KeyType" xml:"KeyType"`
	TaskType              string         `json:"TaskType" xml:"TaskType"`
	Accounts              []AccountsItem `json:"Accounts" xml:"Accounts"`
}
