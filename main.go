/*
Copyright © 2024 c2dp <futurenear@163.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"fmt"
	"github.com/c2dp/zenkit/internal"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

var (
	ctx = gctx.New()
)

func main() {
	//cmd.Execute()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "源地址", "目标地址", "PacketLoss", "AvgRtt"})

	cfg, err := gcfg.New()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	addrSet, err := cfg.Get(ctx, "ping")
	if err != nil {
		g.Log().Error(ctx, err)

	}
	for idx, addr := range addrSet.Strings() {

		receive, err := internal.Ping(3, addr)
		if err != nil {
			g.Log().Errorf(ctx, "ping: %v", err)
		}
		t.AppendRows([]table.Row{
			{idx, receive.SrcIpAddr, receive.DestIpAddr, fmt.Sprintf("%.2f%%", receive.PacketLoss*100), fmt.Sprintf("%.3fms", receive.AvgRtt)},
		})
	}
	t.AppendSeparator()
	t.Render()
}
