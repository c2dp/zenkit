package internal

import (
	"fmt"
	"github.com/c2dp/zenkit/pkg"
	"github.com/go-ping/ping"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"os/signal"
	"runtime"
	"time"
)

type PingReceive struct {
	PacketLoss float64
	SrcIpAddr  string
	DestIpAddr string
	AvgRtt     float64
}

var (
	ctx = gctx.New()
)

func singlePing(cnt int, src, desc string) (*PingReceive, error) {
	pinger, err := ping.NewPinger(desc)
	if err != nil {
		return nil, err
	}
	pinger.Source = src
	if runtime.GOOS != "linux" || os.Getuid() == 0 {
		pinger.SetPrivileged(true)

	}
	pinger.Count = cnt
	pinger.Timeout = 3 * time.Second
	// Listen for Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		g.Log().Debugf(ctx, "%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		g.Log().Debugf(ctx, "%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		g.Log().Debugf(ctx, "\n--- %s ping statistics ---\n", stats.Addr)
		g.Log().Debugf(ctx, "%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		g.Log().Debugf(ctx, "round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	g.Log().Debugf(ctx, "PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	stats := pinger.Statistics()
	return &PingReceive{
		PacketLoss: stats.PacketLoss,
		SrcIpAddr:  pinger.Source,
		DestIpAddr: stats.Addr,
		AvgRtt:     float64(stats.AvgRtt / 1e6),
	}, nil
}
func BatchPing() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"序号", "源地址", "目标地址", "丢包率", "平均往返时间"})

	addrSet, err := g.Cfg().Get(ctx, "ping.descIP")
	if err != nil {
		g.Log().Error(ctx, err)

	}

	addrs, err := pkg.GetLocalAddr()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	for _, src := range addrs {

		for idx, addr := range addrSet.Strings() {

			receive, err := singlePing(3, src.String(), addr)
			if err != nil {
				g.Log().Errorf(ctx, "pingsss: %v", err)
			}
			t.AppendRows([]table.Row{
				{
					idx,
					receive.SrcIpAddr,
					receive.DestIpAddr,
					fmt.Sprintf("%4.0f%%", receive.PacketLoss),
					fmt.Sprintf("%8.3fms", receive.AvgRtt),
				},
			})
		}
	}
	t.AppendSeparator()
	t.Render()
}
