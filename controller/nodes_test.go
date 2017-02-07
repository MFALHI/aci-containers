// Copyright 2017 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"net"
	"testing"
	"time"

	"github.com/noironetworks/aci-containers/ipam"
	"github.com/noironetworks/aci-containers/metadata"
	tu "github.com/noironetworks/aci-containers/testutil"
)

func waitForSEpAnnot(t *testing.T, cont *testAciController, ipv4 net.IP, ipv6 net.IP, desc string) {
	tu.WaitFor(t, desc, 500*time.Millisecond,
		func(last bool) (bool, error) {
			if !tu.WaitCondition(t, last, func() bool {
				return len(cont.nodeUpdates) >= 1
			}, desc, "update") {
				return false, nil
			}

			annot := cont.nodeUpdates[len(cont.nodeUpdates)-1].
				ObjectMeta.Annotations[metadata.ServiceEpAnnotation]

			ep := &metadata.ServiceEndpoint{}
			err := json.Unmarshal([]byte(annot), ep)
			if !tu.WaitNil(t, last, err, desc, "unmarshal", err) {
				return false, nil
			}
			_, err = net.ParseMAC(ep.Mac)
			return tu.WaitNil(t, last, err, desc, "hardware addr parse", err) &&
				tu.WaitEqual(t, last, ipv4, ep.Ipv4, desc, "ipv4") &&
				tu.WaitEqual(t, last, ipv6, ep.Ipv6, desc, "ipv6"), nil
		})
}

func TestServiceEpAnnotationV4(t *testing.T) {
	cont := testController()
	cont.config.NodeServiceIpPool = []ipam.IpRange{
		ipam.IpRange{net.ParseIP("10.1.1.2"), net.ParseIP("10.1.1.3")},
	}
	cont.aciController.initIpam()
	cont.run()

	cont.fakeNodeSource.Add(node("node1"))
	waitForSEpAnnot(t, cont, net.ParseIP("10.1.1.2"), nil, "simple")

	cont.nodeUpdates = nil
	cont.fakeNodeSource.Add(node("node2"))
	waitForSEpAnnot(t, cont, net.ParseIP("10.1.1.3"), nil, "second")

	cont.nodeUpdates = nil
	cont.fakeNodeSource.Add(node("node3"))
	waitForSEpAnnot(t, cont, nil, nil, "noneleft")

	cont.stop()
}

func TestServiceEpAnnotationV6(t *testing.T) {
	cont := testController()
	cont.config.NodeServiceIpPool = []ipam.IpRange{
		ipam.IpRange{net.ParseIP("fd43:85d7:bcf2:9ad2::2"), net.ParseIP("fd43:85d7:bcf2:9ad2::3")},
	}
	cont.aciController.initIpam()
	cont.run()

	cont.fakeNodeSource.Add(node("node1"))
	waitForSEpAnnot(t, cont, nil, net.ParseIP("fd43:85d7:bcf2:9ad2::2"), "simple")

	cont.nodeUpdates = nil
	cont.fakeNodeSource.Add(node("node2"))
	waitForSEpAnnot(t, cont, nil, net.ParseIP("fd43:85d7:bcf2:9ad2::3"), "second")

	cont.nodeUpdates = nil
	cont.fakeNodeSource.Add(node("node3"))
	waitForSEpAnnot(t, cont, nil, nil, "noneleft")

	cont.stop()
}

func TestServiceEpAnnotationExisting(t *testing.T) {
	cont := testController()
	cont.config.NodeServiceIpPool = []ipam.IpRange{
		ipam.IpRange{net.ParseIP("10.1.1.2"), net.ParseIP("10.1.1.4")},
		ipam.IpRange{net.ParseIP("fd43:85d7:bcf2:9ad2::2"), net.ParseIP("fd43:85d7:bcf2:9ad2::4")},
	}
	cont.aciController.initIpam()
	cont.run()

	ep := &metadata.ServiceEndpoint{
		Ipv4: net.ParseIP("10.1.1.1"),
		Ipv6: net.ParseIP("fd43:85d7:bcf2:9ad2::1"),
	}
	n := node("node1")
	raw, _ := json.Marshal(ep)
	n.ObjectMeta.Annotations[metadata.ServiceEpAnnotation] = string(raw)
	cont.fakeNodeSource.Add(n)
	waitForSEpAnnot(t, cont, net.ParseIP("10.1.1.2"),
		net.ParseIP("fd43:85d7:bcf2:9ad2::2"), "out of range")

	cont.nodeUpdates = nil
	n = node("node2")
	ep.Mac = "00:0c:29:92:fe:d0"
	ep.Ipv4 = net.ParseIP("10.1.1.4")
	raw, _ = json.Marshal(ep)
	n.ObjectMeta.Annotations[metadata.ServiceEpAnnotation] = string(raw)
	cont.fakeNodeSource.Add(n)
	waitForSEpAnnot(t, cont, net.ParseIP("10.1.1.4"),
		net.ParseIP("fd43:85d7:bcf2:9ad2::3"), "in range")

	cont.stop()
}
