package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"routecheck/routes"
)

func main() {
	var nets = routes.GetNets()

	for i1 := 0; i1 < len(nets); i1++ {
		for i2 := 0; i2 < len(nets); i2++ {
			if i1 == i2 {
				continue
			}

			if reflect.DeepEqual(nets[i1].IP, net.IP([]byte{0, 0, 0, 0})) {
				continue
			}
			if reflect.DeepEqual(nets[i2].IP, net.IP([]byte{0, 0, 0, 0})) {
				continue
			}

			if reflect.DeepEqual(nets[i1].Mask, net.IPMask([]byte{255, 255, 255, 255})) {
				continue
			}
			if reflect.DeepEqual(nets[i2].Mask, net.IPMask([]byte{255, 255, 255, 255})) {
				continue
			}

			if reflect.DeepEqual(nets[i1].IP, nets[i2].IP) {
				continue
			}

			if netsIntersect(nets[i1], nets[i2]) {
				fmt.Println("network intersection between ", nets[i1], " and ", nets[i2])
				os.Exit(1)
			}
		}
	}
}

func netsIntersect(n1, n2 net.IPNet) bool {
	return n1.Contains(n2.IP) || n2.Contains(n1.IP)
}
