package alidns

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

// DescribeTransferDomains invokes the alidns.DescribeTransferDomains API synchronously
// api document: https://help.aliyun.com/api/alidns/describetransferdomains.html
func (client *Client) DescribeTransferDomains(request *DescribeTransferDomainsRequest) (response *DescribeTransferDomainsResponse, err error) {
	response = CreateDescribeTransferDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTransferDomainsWithChan invokes the alidns.DescribeTransferDomains API asynchronously
// api document: https://help.aliyun.com/api/alidns/describetransferdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTransferDomainsWithChan(request *DescribeTransferDomainsRequest) (<-chan *DescribeTransferDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeTransferDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTransferDomains(request)
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

// DescribeTransferDomainsWithCallback invokes the alidns.DescribeTransferDomains API asynchronously
// api document: https://help.aliyun.com/api/alidns/describetransferdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTransferDomainsWithCallback(request *DescribeTransferDomainsRequest, callback func(response *DescribeTransferDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTransferDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTransferDomains(request)
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

// DescribeTransferDomainsRequest is the request struct for api DescribeTransferDomains
type DescribeTransferDomainsRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	TransferType string           `position:"Query" name:"TransferType"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DescribeTransferDomainsResponse is the response struct for api DescribeTransferDomains
type DescribeTransferDomainsResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	TotalCount      int64           `json:"TotalCount" xml:"TotalCount"`
	PageNumber      int64           `json:"PageNumber" xml:"PageNumber"`
	PageSize        int64           `json:"PageSize" xml:"PageSize"`
	DomainTransfers DomainTransfers `json:"DomainTransfers" xml:"DomainTransfers"`
}

// CreateDescribeTransferDomainsRequest creates a request to invoke DescribeTransferDomains API
func CreateDescribeTransferDomainsRequest() (request *DescribeTransferDomainsRequest) {
	request = &DescribeTransferDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeTransferDomains", "alidns", "openAPI")
	return
}

// CreateDescribeTransferDomainsResponse creates a response to parse from DescribeTransferDomains response
func CreateDescribeTransferDomainsResponse() (response *DescribeTransferDomainsResponse) {
	response = &DescribeTransferDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
