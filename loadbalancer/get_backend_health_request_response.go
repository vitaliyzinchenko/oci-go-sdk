// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetBackendHealthRequest wrapper for the GetBackendHealth operation
type GetBackendHealthRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend server health status to be retrieved.
	LoadBalancerID *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the backend set associated with the backend server to retrieve the health status for.
	// Example: `My_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The IP address and port of the backend server to retrieve the health status for.
	// Example: `1.1.1.7:42`
	BackendName *string `mandatory:"true" contributesTo:"path" name:"backendName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetBackendHealthRequest) String() string {
	return common.PointerString(request)
}

// GetBackendHealthResponse wrapper for the GetBackendHealth operation
type GetBackendHealthResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BackendHealth instance
	BackendHealth `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetBackendHealthResponse) String() string {
	return common.PointerString(response)
}