// The MIT License (MIT)
//
// Copyright (c) 2013-2015 Oryx(ossrs)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package agent
import(
	"fmt"
	"github.com/ossrs/go-oryx/core"
	"github.com/ossrs/go-oryx/protocol"
	"net"
)

type HLS struct {
	endpoint string
	wc       core.WorkerContainer
	l        net.Listener
}

func NewHLS(wc core.WorkerContainer) (agent core.OpenCloser){
	fmt.Println("new HLS Initialize")
	v := &HLS{
		wc:wc,
	}
	core.Conf.Subscribe(v)
	return v
}

// interface core.Agent
func (v *HLS) Open() (err error){
	return
}

func (v *HLS) Close() (err error){
	return
}

func (v *HLS) close() (err error){
	return
}

func (v *HLS) applyListen (c *core.Config) (err error){
	return
}

func (v *HLS) serve(c net.Conn){
	return
}

func (v *HLS) cycle(conn *protocol.RtmpConnection) (err error){
	return
}

func (v *HLS) OnReloadGlobal(scope int, cc, pc *core.Config) (err error){
	return
}

func (v *HLS) OnReloadVhost(vhost string, scope int,cc,pc *core.Config) (err error){
	return
}
