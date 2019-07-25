package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan/examples"
	"github.com/vmware/govmomi/vsan/methods"
	"github.com/vmware/govmomi/vsan/types"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c, err := examples.NewVSANClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	cluster := vim.ManagedObjectReference{
		Type:  "ClusterComputeResource",
		Value: "domain-c8",
	}

	spaceManager := vim.ManagedObjectReference{
		Type:  "VsanSpaceReportSystem",
		Value: "vsan-cluster-space-report-system",
	}

	resp, err := methods.VsanQuerySpaceUsage(ctx, c,
		&types.VsanQuerySpaceUsage{
			This:    spaceManager,
			Cluster: cluster,
		})

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %+v\n", err)
		log.Fatal(err)
	}

	for i := 0; i < reflect.TypeOf(resp.Returnval).NumField(); i++ {
		field := reflect.TypeOf(resp.Returnval).Field(i).Name
		value := reflect.ValueOf(resp.Returnval).Field(i)
		fmt.Fprintf(os.Stdout, "%s:%v\n", field, value)
	}
}
