package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv6"
)

type link struct {
	s     string
	a     netip.Addr
	drv   string
	speed string
	pci   string
}

var links []link
var seq int
var completed = make(map[[2]string]struct{})

func checkConn(link0, link1 link) {
	iface0, iface1 := link0.s, link1.s
	addr0, addr1 := link0.a, link1.a

	var key [2]string
	if strings.Compare(iface0, iface1) < 0 {
		key[0], key[1] = iface0, iface1
	} else {
		key[0], key[1] = iface1, iface0
	}
	if _, ok := completed[key]; ok {
		return
	}

	c, err := icmp.ListenPacket("udp6", addr0.String()+"%"+iface0)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	seq++
	wm := icmp.Message{
		Type: ipv6.ICMPTypeEchoRequest, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: seq,
			Data: []byte("Loopy?"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	dest := &net.UDPAddr{IP: addr1.AsSlice(), Zone: iface0}
	if _, err := c.WriteTo(wb, dest); err != nil {
		log.Fatal(err)
	}
	c.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	var buf [1024]byte
	sz, _, err := c.ReadFrom(buf[:])
	if errors.Is(err, os.ErrDeadlineExceeded) {
		return
	}
	rm, err := icmp.ParseMessage(58, buf[:sz])
	if err != nil {
		return
	}
	if rm.Type != ipv6.ICMPTypeEchoReply {
		return
	}

	completed[key] = struct{}{}
	fmt.Printf("\t%-14s %-10s | %-14s %-10s | %8s\n", iface0, link0.drv, iface1, link1.drv, link0.speed)
}

func getDriverName(iface string) string {
	link, err := os.Readlink("/sys/class/net/" + iface + "/device/driver")
	if err != nil {
		return "???"
	}
	return filepath.Base(link)
}

func getSpeed(iface string) string {
	speedStr, err := os.ReadFile("/sys/class/net/" + iface + "/speed")
	if err != nil {
		return "???"
	}
	speed, err := strconv.ParseUint(strings.TrimSpace(string(speedStr)), 10, 32)
	if err != nil {
		return "???"
	}
	return strconv.FormatFloat(float64(speed)/1000, 'f', 0, 64) + " Gbps"
}

func getPciPath(iface string) string {
	link, err := os.Readlink("/sys/class/net/" + iface + "/device")
	if err != nil {
		return "???"
	}
	return filepath.Base(link)
}

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	sort.Slice(interfaces, func(i, j int) bool {
		return natComp(interfaces[i].Name, interfaces[j].Name)
	})

	links = make([]link, 0, len(interfaces))

	fmt.Println()
	fmt.Println("Addresses:")
	for _, iface := range interfaces {
		pci := getPciPath(iface.Name)
		driver := getDriverName(iface.Name)
		speed := getSpeed(iface.Name)

		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, vAddr := range addrs {
			var addr netip.Addr
			var ok bool
			switch v := vAddr.(type) {
			case *net.IPAddr:
				addr, ok = netip.AddrFromSlice(v.IP)
			case *net.IPNet:
				addr, ok = netip.AddrFromSlice(v.IP)
			}
			if ok && addr.Is6() && addr.IsLinkLocalUnicast() {
				fmt.Printf("\t%-14s %s  %s  %s  %-10s\n", iface.Name, addr.StringExpanded(), iface.HardwareAddr, pci, driver)
				links = append(links, link{iface.Name, addr, driver, speed, pci})
			}
		}
	}
	fmt.Println()

	fmt.Println("Connectivity:")
	for i, link0 := range links {
		for j, link1 := range links {
			if i != j {
				checkConn(link0, link1)
			}
		}
	}
	fmt.Println()
}
