package nat

import (
	"context"
	"testing"
	"time"
)

func TestNat(t *testing.T) {
	port := 1935
	natGateway, err := DiscoverGateway(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log("NAT网关类型:", natGateway.Type())

	natDeviceAddress, err := natGateway.GetDeviceAddress()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("NAT网关IP:", natDeviceAddress.String())

	natExternalAddress, err := natGateway.GetExternalAddress()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("NAT网关公网IP:", natExternalAddress.String())

	ctx := context.Background()
	//进行端口映射
	mappedExternalPort, e := natGateway.AddPortMapping(ctx, "udp", port, "P2P测试", time.Minute)
	if e != nil {
		t.Fatal(e)
	}
	t.Log("NAT映射端口:", port, mappedExternalPort)

	//移除端口映射
	_ = natGateway.DeletePortMapping(ctx, "udp", port)
}
