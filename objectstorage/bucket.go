// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
/* Common set of Object and Archive Storage APIs for managing buckets and objects.
 */

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Bucket A bucket is a container for storing objects in a compartment within a namespace. A bucket is associated with a single compartment.
// The compartment has policies that indicate what actions a user can perform on a bucket and all the objects in the bucket. For more
// information, see [Managing Buckets]({{DOC_SERVER_URL}}/Content/Object/Tasks/managingbuckets.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Bucket struct {

	// The namespace in which the bucket lives.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: my-new-bucket1
	Name *string `mandatory:"true" json:"name"`

	// The compartment ID in which the bucket is authorized.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Arbitrary string keys and values for user-defined metadata.
	Metadata map[string]string `mandatory:"true" json:"metadata"`

	// The OCID of the user who created the bucket.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/rfc/rfc2616), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The entity tag for the bucket.
	Etag *string `mandatory:"true" json:"etag"`

	// The type of public access enabled on this bucket.
	// A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the
	// bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the
	// `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the
	// bucket, public access is allowed for the `GetObject` and `HeadObject` operations.
	PublicAccessType BucketPublicAccessTypeEnum `mandatory:"false" json:"publicAccessType,omitempty"`

	// The type of storage tier of this bucket.
	// A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier.
	// When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier'
	// property is immutable once the bucket is created.
	StorageTier BucketStorageTierEnum `mandatory:"false" json:"storageTier,omitempty"`
}

func (m Bucket) String() string {
	return common.PointerString(m)
}

// BucketPublicAccessTypeEnum Enum with underlying type: string
type BucketPublicAccessTypeEnum string

// Set of constants representing the allowable values for BucketPublicAccessType
const (
	BucketPublicAccessTypeNopublicaccess        BucketPublicAccessTypeEnum = "NoPublicAccess"
	BucketPublicAccessTypeObjectread            BucketPublicAccessTypeEnum = "ObjectRead"
	BucketPublicAccessTypeObjectreadwithoutlist BucketPublicAccessTypeEnum = "ObjectReadWithoutList"
	BucketPublicAccessTypeUnknown               BucketPublicAccessTypeEnum = "UNKNOWN"
)

var mappingBucketPublicAccessType = map[string]BucketPublicAccessTypeEnum{
	"NoPublicAccess":        BucketPublicAccessTypeNopublicaccess,
	"ObjectRead":            BucketPublicAccessTypeObjectread,
	"ObjectReadWithoutList": BucketPublicAccessTypeObjectreadwithoutlist,
	"UNKNOWN":               BucketPublicAccessTypeUnknown,
}

// GetBucketPublicAccessTypeEnumValues Enumerates the set of values for BucketPublicAccessType
func GetBucketPublicAccessTypeEnumValues() []BucketPublicAccessTypeEnum {
	values := make([]BucketPublicAccessTypeEnum, 0)
	for _, v := range mappingBucketPublicAccessType {
		if v != BucketPublicAccessTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}

// BucketStorageTierEnum Enum with underlying type: string
type BucketStorageTierEnum string

// Set of constants representing the allowable values for BucketStorageTier
const (
	BucketStorageTierStandard BucketStorageTierEnum = "Standard"
	BucketStorageTierArchive  BucketStorageTierEnum = "Archive"
	BucketStorageTierUnknown  BucketStorageTierEnum = "UNKNOWN"
)

var mappingBucketStorageTier = map[string]BucketStorageTierEnum{
	"Standard": BucketStorageTierStandard,
	"Archive":  BucketStorageTierArchive,
	"UNKNOWN":  BucketStorageTierUnknown,
}

// GetBucketStorageTierEnumValues Enumerates the set of values for BucketStorageTier
func GetBucketStorageTierEnumValues() []BucketStorageTierEnum {
	values := make([]BucketStorageTierEnum, 0)
	for _, v := range mappingBucketStorageTier {
		if v != BucketStorageTierUnknown {
			values = append(values, v)
		}
	}
	return values
}
