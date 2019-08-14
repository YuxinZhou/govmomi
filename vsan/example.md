Run vSAN examples:
```
go run  ~/govmomi/vsan/examples/health/health.go --url https://<user>:<password>@<vcenter-ip>/sdk
go run  ~/govmomi/vsan/examples/cns/cns.go --url https://<user>:<password>@<vcenter-ip>/sdk
go run  ~/govmomi/vsan/examples/space/space.go --url https://<user>:<password>@<vcenter-ip>/sdk
```

Notes:
1. you might need to escape special characters in the username or password with a backslash (\)
2. you might need to change the moid of the cluster in the code:

```go
cluster := vim.ManagedObjectReference{
    Type:  "ClusterComputeResource",
    Value: "domain-c8", // vSAN MOID
}
```

