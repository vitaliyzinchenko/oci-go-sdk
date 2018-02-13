// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LaunchOptions Options for tuning compatibility and performance of VM shapes.
type LaunchOptions struct {

	// Emulation type for volume.
	// * `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block
	// Storage volumes on Oracle provided images.
	// * `SCSI` - Emulated SCSI disk.
	// * `IDE` - Emulated IDE disk.
	// * `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data
	// volumes on Oracle provided images.
	BootVolumeType LaunchOptionsBootVolumeTypeEnum `mandatory:"true" json:"bootVolumeType"`

	// Firmware used to boot VM.  Select the option that matches your operating system.
	// * `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating
	// systems that boot using MBR style bootloaders.
	// * `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the
	// default for Oracle provided images.
	Firmware LaunchOptionsFirmwareEnum `mandatory:"true" json:"firmware"`

	// Emulation type for NIC.
	// * `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
	// * `VFIO` - Direct attached Virtual Function network controller.  Default for Oracle provided images.
	NetworkType LaunchOptionsNetworkTypeEnum `mandatory:"true" json:"networkType"`

	// Emulation type for volume.
	// * `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block
	// Storage volumes on Oracle provided images.
	// * `SCSI` - Emulated SCSI disk.
	// * `IDE` - Emulated IDE disk.
	// * `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data
	// volumes on Oracle provided images.
	RemoteDataVolumeType LaunchOptionsRemoteDataVolumeTypeEnum `mandatory:"true" json:"remoteDataVolumeType"`
}

func (m LaunchOptions) String() string {
	return common.PointerString(m)
}

// LaunchOptionsBootVolumeTypeEnum Enum with underlying type: string
type LaunchOptionsBootVolumeTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsBootVolumeType
const (
	LaunchOptionsBootVolumeTypeIscsi   LaunchOptionsBootVolumeTypeEnum = "ISCSI"
	LaunchOptionsBootVolumeTypeScsi    LaunchOptionsBootVolumeTypeEnum = "SCSI"
	LaunchOptionsBootVolumeTypeIde     LaunchOptionsBootVolumeTypeEnum = "IDE"
	LaunchOptionsBootVolumeTypeVfio    LaunchOptionsBootVolumeTypeEnum = "VFIO"
	LaunchOptionsBootVolumeTypeUnknown LaunchOptionsBootVolumeTypeEnum = "UNKNOWN"
)

var mappingLaunchOptionsBootVolumeType = map[string]LaunchOptionsBootVolumeTypeEnum{
	"ISCSI":   LaunchOptionsBootVolumeTypeIscsi,
	"SCSI":    LaunchOptionsBootVolumeTypeScsi,
	"IDE":     LaunchOptionsBootVolumeTypeIde,
	"VFIO":    LaunchOptionsBootVolumeTypeVfio,
	"UNKNOWN": LaunchOptionsBootVolumeTypeUnknown,
}

// GetLaunchOptionsBootVolumeTypeEnumValues Enumerates the set of values for LaunchOptionsBootVolumeType
func GetLaunchOptionsBootVolumeTypeEnumValues() []LaunchOptionsBootVolumeTypeEnum {
	values := make([]LaunchOptionsBootVolumeTypeEnum, 0)
	for _, v := range mappingLaunchOptionsBootVolumeType {
		if v != LaunchOptionsBootVolumeTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}

// LaunchOptionsFirmwareEnum Enum with underlying type: string
type LaunchOptionsFirmwareEnum string

// Set of constants representing the allowable values for LaunchOptionsFirmware
const (
	LaunchOptionsFirmwareBios    LaunchOptionsFirmwareEnum = "BIOS"
	LaunchOptionsFirmwareUefi64  LaunchOptionsFirmwareEnum = "UEFI_64"
	LaunchOptionsFirmwareUnknown LaunchOptionsFirmwareEnum = "UNKNOWN"
)

var mappingLaunchOptionsFirmware = map[string]LaunchOptionsFirmwareEnum{
	"BIOS":    LaunchOptionsFirmwareBios,
	"UEFI_64": LaunchOptionsFirmwareUefi64,
	"UNKNOWN": LaunchOptionsFirmwareUnknown,
}

// GetLaunchOptionsFirmwareEnumValues Enumerates the set of values for LaunchOptionsFirmware
func GetLaunchOptionsFirmwareEnumValues() []LaunchOptionsFirmwareEnum {
	values := make([]LaunchOptionsFirmwareEnum, 0)
	for _, v := range mappingLaunchOptionsFirmware {
		if v != LaunchOptionsFirmwareUnknown {
			values = append(values, v)
		}
	}
	return values
}

// LaunchOptionsNetworkTypeEnum Enum with underlying type: string
type LaunchOptionsNetworkTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsNetworkType
const (
	LaunchOptionsNetworkTypeE1000   LaunchOptionsNetworkTypeEnum = "E1000"
	LaunchOptionsNetworkTypeVfio    LaunchOptionsNetworkTypeEnum = "VFIO"
	LaunchOptionsNetworkTypeUnknown LaunchOptionsNetworkTypeEnum = "UNKNOWN"
)

var mappingLaunchOptionsNetworkType = map[string]LaunchOptionsNetworkTypeEnum{
	"E1000":   LaunchOptionsNetworkTypeE1000,
	"VFIO":    LaunchOptionsNetworkTypeVfio,
	"UNKNOWN": LaunchOptionsNetworkTypeUnknown,
}

// GetLaunchOptionsNetworkTypeEnumValues Enumerates the set of values for LaunchOptionsNetworkType
func GetLaunchOptionsNetworkTypeEnumValues() []LaunchOptionsNetworkTypeEnum {
	values := make([]LaunchOptionsNetworkTypeEnum, 0)
	for _, v := range mappingLaunchOptionsNetworkType {
		if v != LaunchOptionsNetworkTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}

// LaunchOptionsRemoteDataVolumeTypeEnum Enum with underlying type: string
type LaunchOptionsRemoteDataVolumeTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsRemoteDataVolumeType
const (
	LaunchOptionsRemoteDataVolumeTypeIscsi   LaunchOptionsRemoteDataVolumeTypeEnum = "ISCSI"
	LaunchOptionsRemoteDataVolumeTypeScsi    LaunchOptionsRemoteDataVolumeTypeEnum = "SCSI"
	LaunchOptionsRemoteDataVolumeTypeIde     LaunchOptionsRemoteDataVolumeTypeEnum = "IDE"
	LaunchOptionsRemoteDataVolumeTypeVfio    LaunchOptionsRemoteDataVolumeTypeEnum = "VFIO"
	LaunchOptionsRemoteDataVolumeTypeUnknown LaunchOptionsRemoteDataVolumeTypeEnum = "UNKNOWN"
)

var mappingLaunchOptionsRemoteDataVolumeType = map[string]LaunchOptionsRemoteDataVolumeTypeEnum{
	"ISCSI":   LaunchOptionsRemoteDataVolumeTypeIscsi,
	"SCSI":    LaunchOptionsRemoteDataVolumeTypeScsi,
	"IDE":     LaunchOptionsRemoteDataVolumeTypeIde,
	"VFIO":    LaunchOptionsRemoteDataVolumeTypeVfio,
	"UNKNOWN": LaunchOptionsRemoteDataVolumeTypeUnknown,
}

// GetLaunchOptionsRemoteDataVolumeTypeEnumValues Enumerates the set of values for LaunchOptionsRemoteDataVolumeType
func GetLaunchOptionsRemoteDataVolumeTypeEnumValues() []LaunchOptionsRemoteDataVolumeTypeEnum {
	values := make([]LaunchOptionsRemoteDataVolumeTypeEnum, 0)
	for _, v := range mappingLaunchOptionsRemoteDataVolumeType {
		if v != LaunchOptionsRemoteDataVolumeTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}
