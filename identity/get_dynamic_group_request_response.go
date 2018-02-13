// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDynamicGroupRequest wrapper for the GetDynamicGroup operation
type GetDynamicGroupRequest struct {

	// The OCID of the dynamic group.
	DynamicGroupId *string `mandatory:"true" contributesTo:"path" name:"dynamicGroupId"`
}

func (request GetDynamicGroupRequest) String() string {
	return common.PointerString(request)
}

// GetDynamicGroupResponse wrapper for the GetDynamicGroup operation
type GetDynamicGroupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DynamicGroup instance
	DynamicGroup `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetDynamicGroupResponse) String() string {
	return common.PointerString(response)
}
