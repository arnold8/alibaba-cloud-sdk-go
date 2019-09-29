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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateStaticVerification invokes the r_kvstore.CreateStaticVerification API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/createstaticverification.html
func (client *Client) CreateStaticVerification(request *CreateStaticVerificationRequest) (response *CreateStaticVerificationResponse, err error) {
	response = CreateCreateStaticVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStaticVerificationWithChan invokes the r_kvstore.CreateStaticVerification API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/createstaticverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStaticVerificationWithChan(request *CreateStaticVerificationRequest) (<-chan *CreateStaticVerificationResponse, <-chan error) {
	responseChan := make(chan *CreateStaticVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStaticVerification(request)
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

// CreateStaticVerificationWithCallback invokes the r_kvstore.CreateStaticVerification API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/createstaticverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStaticVerificationWithCallback(request *CreateStaticVerificationRequest, callback func(response *CreateStaticVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStaticVerificationResponse
		var err error
		defer close(result)
		response, err = client.CreateStaticVerification(request)
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

// CreateStaticVerificationRequest is the request struct for api CreateStaticVerification
type CreateStaticVerificationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	ReplicaId             string           `position:"Query" name:"ReplicaId"`
	SourceInstanceId      string           `position:"Query" name:"SourceInstanceId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	DestinationInstanceId string           `position:"Query" name:"DestinationInstanceId"`
}

// CreateStaticVerificationResponse is the response struct for api CreateStaticVerification
type CreateStaticVerificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateStaticVerificationRequest creates a request to invoke CreateStaticVerification API
func CreateCreateStaticVerificationRequest() (request *CreateStaticVerificationRequest) {
	request = &CreateStaticVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateStaticVerification", "", "")
	return
}

// CreateCreateStaticVerificationResponse creates a response to parse from CreateStaticVerification response
func CreateCreateStaticVerificationResponse() (response *CreateStaticVerificationResponse) {
	response = &CreateStaticVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
