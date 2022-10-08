package routes

import (
	"github.com/vishvananda/netlink"
	"net"
)

func GetNets() []net.IPNet {
	var nets []net.IPNet

	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		panic(err)
	}

	for _, route := range routes {
		if route.Dst == nil {
			continue
		}
		ipNet := net.IPNet{
			IP:   route.Dst.IP,
			Mask: route.Dst.Mask,
		}
		nets = append(nets, ipNet)
	}

	return nets
}
