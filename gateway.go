package nat

import (
	"net"

	"github.com/darkit/nat/netroute"
)

func getDefaultGateway() (net.IP, error) {
	router, err := netroute.New()
	if err != nil {
		return nil, err
	}

	_, ip, _, err := router.Route(net.IPv4zero)
	return ip, err
}
