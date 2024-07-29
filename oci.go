package main

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

type Instance struct {
	DisplayName string `json:"display_name"`
	SSHPublicKey string
	Shape string `json:"shape"`
	OCPUs int32 `json:"ocpus"`
	MemoryGB int64 `json:"memory_gbs"`
	BootVolume int32 `json:"boot_volume_gbs"`
	Domain string `json:"availability_domain"`
	CompartmentID string `json:"compartment_id"`
	ImageID string `json:"image_id"`
	SubnetID string `json:"subnet_id"`
}

func (i *Instance) claim() error {
	
    configProvider := common.DefaultConfigProvider()

    client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
    if err != nil {
        return err
    }

    createInstanceDetails := core.LaunchInstanceDetails{
        DisplayName:     common.String(i.DisplayName),
		Metadata: map[string]string{"ssh_authorized_keys": i.SSHPublicKey},
        CompartmentId:   common.String(i.CompartmentID),
        Shape:           common.String(i.Shape),
		AvailabilityDomain: common.String(i.Domain),
        SourceDetails: core.InstanceSourceViaImageDetails{
            ImageId:    common.String(i.ImageID),
			BootVolumeSizeInGBs: common.Int64(int64(i.BootVolume)),
        },
        CreateVnicDetails: &core.CreateVnicDetails{
            SubnetId: common.String(i.SubnetID),
        },
		ShapeConfig: &core.LaunchInstanceShapeConfigDetails{
			Ocpus: common.Float32(float32(i.OCPUs)),
			MemoryInGBs: common.Float32(float32(i.MemoryGB)),
		},
    }

    request := core.LaunchInstanceRequest{
        LaunchInstanceDetails: createInstanceDetails,
    }

    response, err := client.LaunchInstance(context.Background(), request)
    if err != nil {
        return err
    }

    fmt.Println("Instance launched with OCID: \n", *response.Instance.Id)

	return nil
}