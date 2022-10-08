package routes

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

func GetNets() []net.IPNet {
	var nets []net.IPNet
	var routes, err = getRoutes()
	if err != nil {
		panic(err)
	}
	for _, route := range routes {
		ipNet := net.IPNet{
			IP:   route.GetForwardDest(),
			Mask: route.GetForwardMask(),
		}
		nets = append(nets, ipNet)

	}

	return nets
}

type TableRow struct {
	ForwardDest      [4]byte
	ForwardMask      [4]byte
	ForwardPolicy    uint32
	ForwardNextHop   [4]byte
	ForwardIfIndex   uint32
	ForwardType      uint32
	ForwardProto     uint32
	ForwardAge       uint32
	ForwardNextHopAS uint32
	ForwardMetric1   uint32
	ForwardMetric2   uint32
	ForwardMetric3   uint32
	ForwardMetric4   uint32
	ForwardMetric5   uint32
}

func (rr *TableRow) GetForwardDest() net.IP {
	var ip = make(net.IP, 4)
	copy(ip, rr.ForwardDest[:])
	return ip
}

func (rr *TableRow) GetForwardMask() net.IPMask {
	var ip = make(net.IPMask, 4)
	copy(ip, rr.ForwardMask[:])
	return ip
}

func getRoutes() ([]TableRow, error) {
	var bufLen uint32
	var getIpForwardTable = syscall.NewLazyDLL("iphlpapi.dll").NewProc("GetIpForwardTable")
	getIpForwardTable.Call(uintptr(0), uintptr(unsafe.Pointer(&bufLen)), 0)

	var r1 uintptr
	var buf = make([]byte, bufLen)
	r1, _, _ = getIpForwardTable.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&bufLen)), 0)

	if r1 != 0 {
		return nil, fmt.Errorf("call to GetIpForwardTable failed with result value：%v", r1)
	}

	var (
		num     = *(*uint32)(unsafe.Pointer(&buf[0]))
		routes  = make([]TableRow, num)
		sr      = uintptr(unsafe.Pointer(&buf[0])) + unsafe.Sizeof(num)
		rowSize = unsafe.Sizeof(TableRow{})

		expectedBufferSize = int(bufLen)
		actualBufferSize   = int(unsafe.Sizeof(num) + rowSize*uintptr(num))
	)

	if expectedBufferSize < actualBufferSize {
		return nil, fmt.Errorf("buffer exceeded the expected size of %v while having a size of: %v", expectedBufferSize, actualBufferSize)
	}

	for i := 0; i < int(num); i++ {
		routes[i] = *((*TableRow)(unsafe.Pointer(sr + (rowSize * uintptr(i)))))
	}

	return routes, nil
}
