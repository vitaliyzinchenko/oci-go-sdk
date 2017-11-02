// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// IScsiVolumeAttachment. An ISCSI volume attachment.
type IScsiVolumeAttachment struct {

	// The type of volume attachment.
	AttachmentType *string `mandatory:"true" json:"attachmentType,omitempty"`

	// The Availability Domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the volume attachment.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the instance the volume is attached to.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The current state of the volume attachment.
	LifecycleState IScsiVolumeAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the volume was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the volume.
	VolumeID *string `mandatory:"true" json:"volumeId,omitempty"`

	// The volume's iSCSI IP address.
	// Example: `169.254.0.2`
	Ipv4 *string `mandatory:"true" json:"ipv4,omitempty"`

	// The target volume's iSCSI Qualified Name in the format defined by RFC 3720.
	// Example: `iqn.2015-12.us.oracle.com:456b0391-17b8-4122-bbf1-f85fc0bb97d9`
	Iqn *string `mandatory:"true" json:"iqn,omitempty"`

	// The volume's iSCSI port.
	// Example: `3260`
	Port *int `mandatory:"true" json:"port,omitempty"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `My volume attachment`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name.
	// (Also called the "CHAP password".)
	// Example: `d6866c0d-298b-48ba-95af-309b4faux45e`
	ChapSecret *string `mandatory:"false" json:"chapSecret,omitempty"`

	// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.
	// Example: `ocid1.volume.oc1.phx.abyhqljrgvttnlx73nmrwfaux7kcvzfs3s66izvxf2h4lgvyndsdsnoiwr5q`
	ChapUsername *string `mandatory:"false" json:"chapUsername,omitempty"`
}

func (model IScsiVolumeAttachment) String() string {
	return common.PointerString(model)
}

type IScsiVolumeAttachmentLifecycleStateEnum string
type IScsiVolumeAttachmentLifecycleState struct{}

const (
	I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHING IScsiVolumeAttachmentLifecycleStateEnum = "ATTACHING"
	I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHED  IScsiVolumeAttachmentLifecycleStateEnum = "ATTACHED"
	I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHING IScsiVolumeAttachmentLifecycleStateEnum = "DETACHING"
	I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHED  IScsiVolumeAttachmentLifecycleStateEnum = "DETACHED"
	I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN   IScsiVolumeAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mapping_iscsivolumeattachment_lifecycleState = map[string]IScsiVolumeAttachmentLifecycleStateEnum{
	"ATTACHING": I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHING,
	"ATTACHED":  I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHED,
	"DETACHING": I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHING,
	"DETACHED":  I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHED,
	"UNKNOWN":   I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN,
}

func (receiver IScsiVolumeAttachmentLifecycleState) Values() []IScsiVolumeAttachmentLifecycleStateEnum {
	values := make([]IScsiVolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mapping_iscsivolumeattachment_lifecycleState {
		if v != I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver IScsiVolumeAttachmentLifecycleState) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if IScsiVolumeAttachmentLifecycleStateEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver IScsiVolumeAttachmentLifecycleState) From(toBeConverted string) IScsiVolumeAttachmentLifecycleStateEnum {
	if val, ok := mapping_iscsivolumeattachment_lifecycleState[toBeConverted]; ok {
		return val
	}
	return I_SCSI_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN
}