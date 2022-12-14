package main

import (
	"flag"
	"fmt"

	"github.com/hashicorp/memberlist"

	// "net"
	"os"
	"strconv"
	"time"
)

var (
	bindPort = flag.Int("port", 8001, "gossip port")
)

func mai1n() {
	flag.Parse()
	hostname, _ := os.Hostname()
	config := memberlist.DefaultLocalConfig()
	config.Name = hostname + "-" + strconv.Itoa(*bindPort)
	// config := memberlist.DefaultLocalConfig()
	config.BindPort = *bindPort
	config.AdvertisePort = *bindPort

	fmt.Println("config.DisableTcpPings", config.DisableTcpPings)
	fmt.Println("config.IndirectChecks", config.IndirectChecks)
	fmt.Println("config.RetransmitMult", config.RetransmitMult)

	fmt.Println("config.PushPullInterval", config.PushPullInterval)

	fmt.Println("config.ProbeInterval", config.ProbeInterval)

	fmt.Println("config.GossipInterval", config.GossipInterval)
	fmt.Println("config.GossipNodes", config.GossipNodes)

	fmt.Println("config.BindPort", config.BindPort)

	list1, err := memberlist.Create(config)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.
	// 配置种子节点, 这里我直接写死了
	_, err = list1.Join([]string{"127.0.0.1:8001", "127.0.0.1:8002"})
	fmt.Println("err", err)

	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	// Ask for members of the cluster
	for {
		fmt.Println("-------------start--------------")
		for _, member := range list1.Members() {
			fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
		}
		fmt.Println("-------------end--------------")
		time.Sleep(time.Second * 3)

	}

}
func ma3in() {
	ints := make([]int, 10, 10)
	fmt.Println(ints)
	for i := 0; i < 2; i++ {
		ints = append(ints, i)
	}
	slice := ints[:0]
	for i := 0; i < 7; i++ {
		slice = append(slice, i+15)
	}
	slice = slice[:0]
	fmt.Println(ints, slice)
}
