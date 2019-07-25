package main

import (
	"context"
	"fmt"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan/examples"
	"github.com/vmware/govmomi/vsan/methods"
	"github.com/vmware/govmomi/vsan/types"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c, err := examples.NewVSANClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	cnsManager := vim.ManagedObjectReference{
		Type:  "CnsVolumeManager",
		Value: "cns-volume-manager",
	}

	resp2, err := methods.CnsQueryAllVolume(ctx, c,
		&types.CnsQueryAllVolume{
			This:      cnsManager,
			Filter:    types.CnsQueryFilter{},
			Selection: types.CnsQuerySelection{},
		})

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %+v\n", err)
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "res: %+v\n", resp2)
}
