package openanalytics_open

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

// ForbidAutomaticMetaSyncAsIntegrationAccount invokes the openanalytics_open.ForbidAutomaticMetaSyncAsIntegrationAccount API synchronously
func (client *Client) ForbidAutomaticMetaSyncAsIntegrationAccount(request *ForbidAutomaticMetaSyncAsIntegrationAccountRequest) (response *ForbidAutomaticMetaSyncAsIntegrationAccountResponse, err error) {
	response = CreateForbidAutomaticMetaSyncAsIntegrationAccountResponse()
	err = client.DoAction(request, response)
	return
}

// ForbidAutomaticMetaSyncAsIntegrationAccountWithChan invokes the openanalytics_open.ForbidAutomaticMetaSyncAsIntegrationAccount API asynchronously
func (client *Client) ForbidAutomaticMetaSyncAsIntegrationAccountWithChan(request *ForbidAutomaticMetaSyncAsIntegrationAccountRequest) (<-chan *ForbidAutomaticMetaSyncAsIntegrationAccountResponse, <-chan error) {
	responseChan := make(chan *ForbidAutomaticMetaSyncAsIntegrationAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ForbidAutomaticMetaSyncAsIntegrationAccount(request)
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

// ForbidAutomaticMetaSyncAsIntegrationAccountWithCallback invokes the openanalytics_open.ForbidAutomaticMetaSyncAsIntegrationAccount API asynchronously
func (client *Client) ForbidAutomaticMetaSyncAsIntegrationAccountWithCallback(request *ForbidAutomaticMetaSyncAsIntegrationAccountRequest, callback func(response *ForbidAutomaticMetaSyncAsIntegrationAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ForbidAutomaticMetaSyncAsIntegrationAccountResponse
		var err error
		defer close(result)
		response, err = client.ForbidAutomaticMetaSyncAsIntegrationAccount(request)
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

// ForbidAutomaticMetaSyncAsIntegrationAccountRequest is the request struct for api ForbidAutomaticMetaSyncAsIntegrationAccount
type ForbidAutomaticMetaSyncAsIntegrationAccountRequest struct {
	*requests.RpcRequest
}

// ForbidAutomaticMetaSyncAsIntegrationAccountResponse is the response struct for api ForbidAutomaticMetaSyncAsIntegrationAccount
type ForbidAutomaticMetaSyncAsIntegrationAccountResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	HasUpdated bool   `json:"HasUpdated" xml:"HasUpdated"`
}

// CreateForbidAutomaticMetaSyncAsIntegrationAccountRequest creates a request to invoke ForbidAutomaticMetaSyncAsIntegrationAccount API
func CreateForbidAutomaticMetaSyncAsIntegrationAccountRequest() (request *ForbidAutomaticMetaSyncAsIntegrationAccountRequest) {
	request = &ForbidAutomaticMetaSyncAsIntegrationAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "ForbidAutomaticMetaSyncAsIntegrationAccount", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateForbidAutomaticMetaSyncAsIntegrationAccountResponse creates a response to parse from ForbidAutomaticMetaSyncAsIntegrationAccount response
func CreateForbidAutomaticMetaSyncAsIntegrationAccountResponse() (response *ForbidAutomaticMetaSyncAsIntegrationAccountResponse) {
	response = &ForbidAutomaticMetaSyncAsIntegrationAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}