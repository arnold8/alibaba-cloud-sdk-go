package mts

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

// Pipeline is a nested struct in mts response
type Pipeline struct {
	Name          string       `json:"Name" xml:"Name"`
	SpeedLevel    int64        `json:"SpeedLevel" xml:"SpeedLevel"`
	Priority      int          `json:"Priority" xml:"Priority"`
	State         string       `json:"State" xml:"State"`
	Role          string       `json:"Role" xml:"Role"`
	Speed         string       `json:"Speed" xml:"Speed"`
	QuotaAllocate int64        `json:"QuotaAllocate" xml:"QuotaAllocate"`
	Id            string       `json:"Id" xml:"Id"`
	NotifyConfig  NotifyConfig `json:"NotifyConfig" xml:"NotifyConfig"`
}
