package main

import (
	"context"
	"fmt"
	"log"
	"syscall"

	md_netlink "github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go demo1()
	go demo2()
	// 退出
	<-ctx.Done()
	log.Println("程序退出")
}

func demo1() {
	conn, err := md_netlink.Dial(unix.NETLINK_ROUTE, &md_netlink.Config{
		Groups: unix.RTNLGRP_LINK,
	})
	if err != nil {
		log.Printf("Could not read netlink: %v", err)
	}
	log.Printf("start receive message")
	for {
		msgs, err := conn.Receive()
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}
		for _, m := range msgs {
			log.Printf("Receive message: %v", m)
			switch m.Header.Type {
			case unix.RTM_DELADDR:
				log.Printf("Receive message RTM_DELADDR")
			case unix.RTM_NEWADDR:
				log.Printf("Receive message RTM_NEWADDR")
			}

			// 根据linkMsg分析网卡添加、删除或配置变化等事件
		}
	}
}

// https://stackoverflow.com/questions/36347807/how-to-monitor-ip-address-change-using-rtnetlink-socket-in-go-language
func demo2() {
	l, _ := ListenNetlink()

	for {
		msgs, err := l.ReadMsgs()
		if err != nil {
			log.Printf("Could not read netlink2: %v", err)
		}

		for _, m := range msgs {
			log.Printf("Receive message2: %v", m)
			if IsNewAddr(&m) {
				log.Println("New Addr")
			}

			if IsDelAddr(&m) {
				log.Println("Del Addr")
			}
		}
	}
}

type NetlinkListener struct {
	fd int
	sa *syscall.SockaddrNetlink
}

func ListenNetlink() (*NetlinkListener, error) {
	groups := (1 << (syscall.RTNLGRP_LINK - 1)) |
		(1 << (syscall.RTNLGRP_IPV4_IFADDR - 1)) |
		(1 << (syscall.RTNLGRP_IPV6_IFADDR - 1))

	s, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM,
		syscall.NETLINK_ROUTE)
	if err != nil {
		return nil, fmt.Errorf("socket: %s", err)
	}

	saddr := &syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pid:    uint32(0),
		Groups: uint32(groups),
	}

	err = syscall.Bind(s, saddr)
	if err != nil {
		return nil, fmt.Errorf("bind: %s", err)
	}

	return &NetlinkListener{fd: s, sa: saddr}, nil
}

func (l *NetlinkListener) ReadMsgs() ([]syscall.NetlinkMessage, error) {
	defer func() {
		recover()
	}()

	pkt := make([]byte, 2048)

	n, err := syscall.Read(l.fd, pkt)
	if err != nil {
		return nil, fmt.Errorf("read: %s", err)
	}

	msgs, err := syscall.ParseNetlinkMessage(pkt[:n])
	if err != nil {
		return nil, fmt.Errorf("parse: %s", err)
	}

	return msgs, nil
}

func IsNewAddr(msg *syscall.NetlinkMessage) bool {
	if msg.Header.Type == syscall.RTM_NEWADDR {
		return true
	}

	return false
}

func IsDelAddr(msg *syscall.NetlinkMessage) bool {
	if msg.Header.Type == syscall.RTM_DELADDR {
		return true
	}

	return false
}

// rtm_scope is the distance to the destination:
//
// RT_SCOPE_UNIVERSE   global route
// RT_SCOPE_SITE       interior route in the
// local autonomous system
// RT_SCOPE_LINK       route on this link
// RT_SCOPE_HOST       route on the local host
// RT_SCOPE_NOWHERE    destination doesn't exist
//
// The values between RT_SCOPE_UNIVERSE and RT_SCOPE_SITE are
// available to the user.

func IsRelevant(msg *syscall.IfAddrmsg) bool {
	if msg.Scope == syscall.RT_SCOPE_UNIVERSE ||
		msg.Scope == syscall.RT_SCOPE_SITE {
		return true
	}

	return false
}
