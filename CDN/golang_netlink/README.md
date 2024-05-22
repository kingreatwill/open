// https://stackoverflow.com/questions/36347807/how-to-monitor-ip-address-change-using-rtnetlink-socket-in-go-language
	// https://www.coder.work/article/204242#google_vignette

	// Subscribe to receive link events
	// if err := conn.SetSubscribe(netlink.SubscriptionOptions{
	// 	Messages: []netlink.MessageType{
	// 		netlink.RTM_NEWLINK, // New link detected
	// 		netlink.RTM_DELLINK, // Link removed
	// 		netlink.RTM_CHGADDR, // Address change
	// 	},
	// }); err != nil {
	// 	log.Fatalf("Failed to set subscription: %v", err)
	// }

	// "github.com/vishvananda/netlink"
	// 设置过滤条件，只关注网络接口事件
	// filter := &netlink.RouteSubscribe{
	// 	Groups: []uint{
	// 		netlink.RTNLGRP_LINK,
	// 	},
	// }
	// if err := conn.SetSubscribe(filter); err != nil {
	// 	log.Fatalf("Failed to set subscription: %v", err)
	// }

	// // 设置过滤器，只接收网卡相关的事件
	// filter := &netlink.GenericNetlinkFamily{
	// 	Family: unix.NETLINK_ROUTE,
	// 	Groups: netlink.RTNLGRP_LINK,
	// }
	// if err := conn.SetSubscribe(filter); err != nil {
	// 	log.Fatalf("Error setting subscription: %v", err)
	// }

	// 循环读取事件
	// for {
	// 	msgs, err := conn.Receive()
	// 	if err != nil {
	// 		log.Printf("Error receiving message: %v", err)
	// 		continue
	// 	}
	// 	for _, m := range msgs {
	// 		linkMsg, ok := m.Data.(*netlink.LinkMessage)
	// 		if !ok {
	// 			log.Printf("Unexpected message type: %T", m.Data)
	// 			continue
	// 		}
	// 		log.Printf("Link event: %+v", linkMsg)
	// 		// 根据linkMsg分析网卡添加、删除或配置变化等事件
	// 	}
	// }

	// 添加一个接收消息的通道
	// msgs, err := conn.Receive()
	// if err != nil {
	// 	log.Fatalf("Failed to start receiving messages: %v", err)
	// 	return
	// }
	// // 定义一个处理函数来处理接收到的消息
	// go func() {
	// 	for {
	// 		select {
	// 		case msg := <-msgs:
	// 			switch msg.Header.Type {
	// 			case netlink.RTM_NEWLINK, netlink.RTM_DELLINK, netlink.RTM_CHANGELINK:
	// 				link, ok := msg.Data.(*netlink.LinkMessage)
	// 				if !ok {
	// 					log.Printf("Received unexpected message type: %d", msg.Header.Type)
	// 					continue
	// 				}

	// 				// 根据消息类型打印网卡变化信息
	// 				linkType := ""
	// 				switch msg.Header.Type {
	// 				case unix.RTM_NEWLINK:
	// 					linkType = "created"
	// 				case unix.RTM_DELLINK:
	// 					linkType = "deleted"
	// 				case unix.RTM_CHANGELINK: netlink.RTM_CHGADDR
	// 					linkType = "changed"
	// 				}
	// 				fmt.Printf("Interface %s %s\n", link.Attributes.Name, linkType)

	// 				// 这里可以进一步处理网卡变化，比如获取详细的网卡信息等
	// 			default:
	// 				// 忽略其他类型的消息
	// 			}
	// 		}
	// 	}
	// }()