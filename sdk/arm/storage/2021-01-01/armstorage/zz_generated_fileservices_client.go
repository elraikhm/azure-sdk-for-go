// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// FileServicesClient contains the methods for the FileServices group.
// Don't use this type directly, use NewFileServicesClient() instead.
type FileServicesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFileServicesClient creates a new instance of FileServicesClient with the specified values.
func NewFileServicesClient(con *armcore.Connection, subscriptionID string) *FileServicesClient {
	return &FileServicesClient{con: con, subscriptionID: subscriptionID}
}

// GetServiceProperties - Gets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
func (client *FileServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesGetServicePropertiesOptions) (FileServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return FileServicePropertiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FileServicePropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FileServicePropertiesResponse{}, client.getServicePropertiesHandleError(resp)
	}
	return client.getServicePropertiesHandleResponse(resp)
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *FileServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesGetServicePropertiesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{FileServicesName}", url.PathEscape("default"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *FileServicesClient) getServicePropertiesHandleResponse(resp *azcore.Response) (FileServicePropertiesResponse, error) {
	var val *FileServiceProperties
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FileServicePropertiesResponse{}, err
	}
	return FileServicePropertiesResponse{RawResponse: resp.Response, FileServiceProperties: val}, nil
}

// getServicePropertiesHandleError handles the GetServiceProperties error response.
func (client *FileServicesClient) getServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List all file services in storage accounts
func (client *FileServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesListOptions) (FileServiceItemsResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return FileServiceItemsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FileServiceItemsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FileServiceItemsResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *FileServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *FileServicesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FileServicesClient) listHandleResponse(resp *azcore.Response) (FileServiceItemsResponse, error) {
	var val *FileServiceItems
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FileServiceItemsResponse{}, err
	}
	return FileServiceItemsResponse{RawResponse: resp.Response, FileServiceItems: val}, nil
}

// listHandleError handles the List error response.
func (client *FileServicesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SetServiceProperties - Sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
func (client *FileServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters FileServiceProperties, options *FileServicesSetServicePropertiesOptions) (FileServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return FileServicePropertiesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FileServicePropertiesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FileServicePropertiesResponse{}, client.setServicePropertiesHandleError(resp)
	}
	return client.setServicePropertiesHandleResponse(resp)
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *FileServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters FileServiceProperties, options *FileServicesSetServicePropertiesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{FileServicesName}", url.PathEscape("default"))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *FileServicesClient) setServicePropertiesHandleResponse(resp *azcore.Response) (FileServicePropertiesResponse, error) {
	var val *FileServiceProperties
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FileServicePropertiesResponse{}, err
	}
	return FileServicePropertiesResponse{RawResponse: resp.Response, FileServiceProperties: val}, nil
}

// setServicePropertiesHandleError handles the SetServiceProperties error response.
func (client *FileServicesClient) setServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}
