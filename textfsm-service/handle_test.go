package function

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"testing"
)

// TestHandle ensures that Handle executes without error and returns the
// HTTP 200 status code indicating no errors.
func TestHandle(t *testing.T) {
	data := map[string]string{
		"stdin":  "ifconfig",
		"stdout": "cni0: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet 10.42.0.1  netmask 255.255.255.0  broadcast 10.42.0.255\n        inet6 fe80::6433:bcff:fe55:55b4  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 66:33:bc:55:55:b4  txqueuelen 1000  (Ethernet)\n        RX packets 24599592  bytes 3866304438 (3.8 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 25197719  bytes 14474932955 (14.4 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\ndocker0: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1500\n        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255\n        inet6 fe80::42:cfff:fec8:ba0c  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 02:42:cf:c8:ba:0c  txqueuelen 0  (Ethernet)\n        RX packets 131  bytes 9339 (9.3 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 204  bytes 26439 (26.4 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nenp3s0: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1500\n        inet 172.23.0.10  netmask 255.255.0.0  broadcast 172.23.255.255\n        inet6 fe80::c6a8:1dff:fe7e:9389  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether c4:a8:1d:7e:93:89  txqueuelen 1000  (Ethernet)\n        RX packets 4169426  bytes 5090842110 (5.0 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 2976180  bytes 391805447 (391.8 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nenp3s0.400: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1500\n        inet 172.24.127.82  netmask 255.255.0.0  broadcast 172.24.255.255\n        inet6 fe80::c6a8:1dff:fe7e:9389  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether c4:a8:1d:7e:93:89  txqueuelen 1000  (Ethernet)\n        RX packets 9319  bytes 2598971 (2.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 6937  bytes 1311944 (1.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nenp3s0.401: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1500\n        inet 172.25.255.216  netmask 255.255.0.0  broadcast 172.25.255.255\n        inet6 fe80::c6a8:1dff:fe7e:9389  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether c4:a8:1d:7e:93:89  txqueuelen 1000  (Ethernet)\n        RX packets 3281154  bytes 4977720994 (4.9 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 2923032  bytes 386248767 (386.2 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nenp3s0.403: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1500\n        inet 172.27.140.153  netmask 255.255.0.0  broadcast 172.27.255.255\n        inet6 fe80::c6a8:1dff:fe7e:9389  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether c4:a8:1d:7e:93:89  txqueuelen 1000  (Ethernet)\n        RX packets 23016  bytes 3793765 (3.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 15798  bytes 1479851 (1.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nenp4s0: flags=4099\u003cUP,BROADCAST,MULTICAST\u003e  mtu 1500\n        ether 60:a4:4c:37:39:4a  txqueuelen 1000  (Ethernet)\n        RX packets 0  bytes 0 (0.0 B)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 0  bytes 0 (0.0 B)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nflannel.1: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet 10.42.0.0  netmask 255.255.255.255  broadcast 0.0.0.0\n        inet6 fe80::fc81:e6ff:fefa:b58b  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether fe:81:e6:fa:b5:8b  txqueuelen 0  (Ethernet)\n        RX packets 0  bytes 0 (0.0 B)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 0  bytes 0 (0.0 B)\n        TX errors 0  dropped 5 overruns 0  carrier 0  collisions 0\n\nlo: flags=73\u003cUP,LOOPBACK,RUNNING\u003e  mtu 65536\n        inet 127.0.0.1  netmask 255.0.0.0\n        inet6 ::1  prefixlen 128  scopeid 0x10\u003chost\u003e\n        loop  txqueuelen 1000  (Local Loopback)\n        RX packets 5519419  bytes 2140116946 (2.1 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 5519419  bytes 2140116946 (2.1 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth90165165: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::895:92ff:feab:4d2  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 0a:95:92:ab:04:d2  txqueuelen 0  (Ethernet)\n        RX packets 204351  bytes 23506811 (23.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 216398  bytes 59914050 (59.9 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth1818ed55: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::54d5:43ff:fea4:a416  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 56:d5:43:a4:a4:16  txqueuelen 0  (Ethernet)\n        RX packets 244198  bytes 32901695 (32.9 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 253440  bytes 51329931 (51.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth18bd24b7: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::4ce2:59ff:fed8:d2a4  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 4e:e2:59:d8:d2:a4  txqueuelen 0  (Ethernet)\n        RX packets 644755  bytes 191505369 (191.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 743890  bytes 134573247 (134.5 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth18bd636d: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::fb:1cff:feb5:1406  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 02:fb:1c:b5:14:06  txqueuelen 0  (Ethernet)\n        RX packets 1815330  bytes 434558307 (434.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1827848  bytes 765304145 (765.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth1d7e101c: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::f0a1:85ff:fee5:7d82  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether f2:a1:85:e5:7d:82  txqueuelen 0  (Ethernet)\n        RX packets 156048  bytes 25554294 (25.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 168464  bytes 45210395 (45.2 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth1dfed069: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::4482:e6ff:fe73:6e19  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 46:82:e6:73:6e:19  txqueuelen 0  (Ethernet)\n        RX packets 1233762  bytes 259205860 (259.2 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1216768  bytes 187551794 (187.5 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth24ac976f: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::80e3:ccff:fee6:5bd1  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 82:e3:cc:e6:5b:d1  txqueuelen 0  (Ethernet)\n        RX packets 152486  bytes 25732243 (25.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 160107  bytes 38710849 (38.7 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth368d61d6: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::f4a1:f9ff:fe62:ccd0  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether f6:a1:f9:62:cc:d0  txqueuelen 0  (Ethernet)\n        RX packets 1198435  bytes 250013019 (250.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1165600  bytes 175892494 (175.8 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth39e13fcf: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::a899:25ff:fede:f27b  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether aa:99:25:de:f2:7b  txqueuelen 0  (Ethernet)\n        RX packets 1684735  bytes 2370870680 (2.3 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1834677  bytes 751748777 (751.7 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth3bf5d981: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::1431:35ff:fe80:967d  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 16:31:35:80:96:7d  txqueuelen 0  (Ethernet)\n        RX packets 207181  bytes 45420731 (45.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 205021  bytes 77885918 (77.8 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth3e66d193: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::a44c:dff:fef8:8944  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether a6:4c:0d:f8:89:44  txqueuelen 0  (Ethernet)\n        RX packets 8911  bytes 694084 (694.0 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 11379  bytes 1578934 (1.5 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth53780c74: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::5829:33ff:fec9:3686  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 5a:29:33:c9:36:86  txqueuelen 0  (Ethernet)\n        RX packets 33  bytes 2362 (2.3 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 3637  bytes 208126 (208.1 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth56236ae6: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::fc42:afff:feec:4d46  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether fe:42:af:ec:4d:46  txqueuelen 0  (Ethernet)\n        RX packets 393128  bytes 39484470 (39.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 436870  bytes 52365857 (52.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth5ce9bca5: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::fc59:ebff:fe6b:b631  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether fe:59:eb:6b:b6:31  txqueuelen 0  (Ethernet)\n        RX packets 23609  bytes 1943515 (1.9 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 32132  bytes 121430122 (121.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth60a0019b: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::f002:fcff:fea2:3cb  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether f2:02:fc:a2:03:cb  txqueuelen 0  (Ethernet)\n        RX packets 399956  bytes 2280551008 (2.2 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 400607  bytes 430443489 (430.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth614c84b: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1500\n        inet6 fe80::24e8:43ff:fe97:2dd2  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 26:e8:43:97:2d:d2  txqueuelen 0  (Ethernet)\n        RX packets 41  bytes 3310 (3.3 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 88  bytes 12404 (12.4 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth647e7633: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::4c21:1ff:fe11:d14  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 4e:21:01:11:0d:14  txqueuelen 0  (Ethernet)\n        RX packets 157655  bytes 12896692 (12.8 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 148805  bytes 13962573 (13.9 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth64c54825: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::e4cc:6cff:feca:882d  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether e6:cc:6c:ca:88:2d  txqueuelen 0  (Ethernet)\n        RX packets 24989  bytes 2073019 (2.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 33159  bytes 122474855 (122.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth77cab68b: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::e0e3:94ff:fe71:ce8  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether e2:e3:94:71:0c:e8  txqueuelen 0  (Ethernet)\n        RX packets 372761  bytes 56779332 (56.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 400299  bytes 43869363 (43.8 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth78a9dded: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::ecda:8aff:fef2:373f  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether ee:da:8a:f2:37:3f  txqueuelen 0  (Ethernet)\n        RX packets 33  bytes 2362 (2.3 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 3210  bytes 176268 (176.2 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth7b63be8f: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::d835:99ff:fe15:d636  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether da:35:99:15:d6:36  txqueuelen 0  (Ethernet)\n        RX packets 53579  bytes 7119201 (7.1 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 58197  bytes 12790555 (12.7 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth7ca88ec8: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::b1:ebff:fe22:d68d  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 02:b1:eb:22:d6:8d  txqueuelen 0  (Ethernet)\n        RX packets 34  bytes 2432 (2.4 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 3443  bytes 193290 (193.2 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth7e3c877a: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::b425:59ff:fefd:9b03  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether b6:25:59:fd:9b:03  txqueuelen 0  (Ethernet)\n        RX packets 566817  bytes 116442173 (116.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 559509  bytes 191975712 (191.9 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth833f9cdc: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::78d7:f1ff:fef5:de76  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 7a:d7:f1:f5:de:76  txqueuelen 0  (Ethernet)\n        RX packets 84117  bytes 11453287 (11.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 87852  bytes 17411072 (17.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth84cbbcc0: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::f07f:daff:fe40:c325  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether f2:7f:da:40:c3:25  txqueuelen 0  (Ethernet)\n        RX packets 1888842  bytes 418459676 (418.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1874738  bytes 454739185 (454.7 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth86a527e7: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::7080:c1ff:fe5a:fa0c  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 72:80:c1:5a:fa:0c  txqueuelen 0  (Ethernet)\n        RX packets 34  bytes 2432 (2.4 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 3332  bytes 184944 (184.9 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth876550f2: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::840:c4ff:fe15:e26d  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 0a:40:c4:15:e2:6d  txqueuelen 0  (Ethernet)\n        RX packets 97617  bytes 20775877 (20.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 97235  bytes 54306264 (54.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth8cc69d62: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::18df:8bff:feeb:5834  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 1a:df:8b:eb:58:34  txqueuelen 0  (Ethernet)\n        RX packets 350981  bytes 30571624 (30.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 396192  bytes 35870129 (35.8 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth8cdddd2d: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::b4f7:36ff:fea9:9a38  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether b6:f7:36:a9:9a:38  txqueuelen 0  (Ethernet)\n        RX packets 59013  bytes 11129237 (11.1 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 57685  bytes 22656517 (22.6 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth8e6ed7cd: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::a8aa:d3ff:fe23:b9f9  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether aa:aa:d3:23:b9:f9  txqueuelen 0  (Ethernet)\n        RX packets 158072  bytes 25888466 (25.8 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 166675  bytes 39099595 (39.0 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth99a95cd4: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::ece5:32ff:fe2e:4020  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether ee:e5:32:2e:40:20  txqueuelen 0  (Ethernet)\n        RX packets 20992  bytes 1950482 (1.9 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 25141  bytes 3249412 (3.2 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth9bf7b94b: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::8075:cdff:fe52:172d  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 82:75:cd:52:17:2d  txqueuelen 0  (Ethernet)\n        RX packets 46330  bytes 4766350 (4.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 52016  bytes 4696064 (4.6 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth9c0984f9: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::60b2:b5ff:fe71:bd1e  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 62:b2:b5:71:bd:1e  txqueuelen 0  (Ethernet)\n        RX packets 1272005  bytes 258248779 (258.2 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1247511  bytes 201193728 (201.1 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth9e04cb63: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::a862:14ff:fe71:117a  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether aa:62:14:71:11:7a  txqueuelen 0  (Ethernet)\n        RX packets 1256620  bytes 274047412 (274.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1281331  bytes 212531634 (212.5 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth9f96f16f: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::3c4e:3cff:feca:cbd1  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 3e:4e:3c:ca:cb:d1  txqueuelen 0  (Ethernet)\n        RX packets 33  bytes 2362 (2.3 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 3696  bytes 212772 (212.7 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nveth9fc03cbd: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::34b3:61ff:fed5:4fae  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 36:b3:61:d5:4f:ae  txqueuelen 0  (Ethernet)\n        RX packets 10096  bytes 877848 (877.8 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 12823  bytes 1893791 (1.8 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvetha18be857: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::2cc3:f4ff:fed1:73aa  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 2e:c3:f4:d1:73:aa  txqueuelen 0  (Ethernet)\n        RX packets 530836  bytes 98419222 (98.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 535954  bytes 160333679 (160.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvetha5b697ae: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::2c58:90ff:fede:76e7  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 2e:58:90:de:76:e7  txqueuelen 0  (Ethernet)\n        RX packets 1263398  bytes 176048398 (176.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1387197  bytes 1656741469 (1.6 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvetha5cc5076: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::6053:11ff:fede:9c2f  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 62:53:11:de:9c:2f  txqueuelen 0  (Ethernet)\n        RX packets 354092  bytes 30805919 (30.8 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 397410  bytes 36195454 (36.1 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvetha8d625d1: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::68e6:beff:fea8:5561  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 6a:e6:be:a8:55:61  txqueuelen 0  (Ethernet)\n        RX packets 1410691  bytes 296004426 (296.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1379093  bytes 256324853 (256.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethaa4f2880: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::f45c:deff:fe59:375e  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether f6:5c:de:59:37:5e  txqueuelen 0  (Ethernet)\n        RX packets 161156  bytes 25450022 (25.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 187625  bytes 588197217 (588.1 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethb2f1469c: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::f053:45ff:fefa:6b13  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether f2:53:45:fa:6b:13  txqueuelen 0  (Ethernet)\n        RX packets 154725  bytes 25787083 (25.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 161420  bytes 38683559 (38.6 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethb3ac8cab: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::3cff:c9ff:fe84:8803  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 3e:ff:c9:84:88:03  txqueuelen 0  (Ethernet)\n        RX packets 8834  bytes 759624 (759.6 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 11009  bytes 1481634 (1.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethbc1e705b: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::a8e1:b6ff:febb:dfc4  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether aa:e1:b6:bb:df:c4  txqueuelen 0  (Ethernet)\n        RX packets 1214036  bytes 101047628 (101.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1455776  bytes 1031982507 (1.0 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethbe16058d: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::18f1:d2ff:fef5:d094  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 1a:f1:d2:f5:d0:94  txqueuelen 0  (Ethernet)\n        RX packets 7948  bytes 603408 (603.4 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 9748  bytes 1445664 (1.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethc08d8f16: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::20ed:78ff:fee3:a2cb  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 22:ed:78:e3:a2:cb  txqueuelen 0  (Ethernet)\n        RX packets 1205697  bytes 253115694 (253.1 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1226652  bytes 191434921 (191.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethc86fac5b: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::4c03:3cff:fee8:a138  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 4e:03:3c:e8:a1:38  txqueuelen 0  (Ethernet)\n        RX packets 148094  bytes 12785177 (12.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 137604  bytes 131694426 (131.6 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethcbc44a8e: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::8cc1:17ff:fe0a:b5f0  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 8e:c1:17:0a:b5:f0  txqueuelen 0  (Ethernet)\n        RX packets 130770  bytes 27689068 (27.6 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 145495  bytes 16008648 (16.0 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethcfdf3eb8: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::a853:7cff:fe75:e5e8  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether aa:53:7c:75:e5:e8  txqueuelen 0  (Ethernet)\n        RX packets 33  bytes 2362 (2.3 KB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 3387  bytes 188978 (188.9 KB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethd48f899e: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::904a:b2ff:fe7c:fb6  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 92:4a:b2:7c:0f:b6  txqueuelen 0  (Ethernet)\n        RX packets 1739035  bytes 156018634 (156.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1472318  bytes 4708544900 (4.7 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethd525468f: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::ac53:d4ff:feed:814  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether ae:53:d4:ed:08:14  txqueuelen 0  (Ethernet)\n        RX packets 108178  bytes 9417036 (9.4 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 128997  bytes 11360921 (11.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethd8eac673: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::3012:90ff:fe90:7e99  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 32:12:90:90:7e:99  txqueuelen 0  (Ethernet)\n        RX packets 72162  bytes 17160475 (17.1 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 74186  bytes 22182321 (22.1 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethe0f51232: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::b03c:f0ff:fe08:e066  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether b2:3c:f0:08:e0:66  txqueuelen 0  (Ethernet)\n        RX packets 1261990  bytes 1219584469 (1.2 GB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1251643  bytes 4489981909 (4.4 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethe0ffbfe9: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::e077:11ff:febd:47a0  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether e2:77:11:bd:47:a0  txqueuelen 0  (Ethernet)\n        RX packets 1078608  bytes 112174318 (112.1 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 1206345  bytes 1545989429 (1.5 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethe3b4cb2e: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::bc06:65ff:fe32:201c  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether be:06:65:32:20:1c  txqueuelen 0  (Ethernet)\n        RX packets 155213  bytes 26013101 (26.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 164091  bytes 39329079 (39.3 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethe5b8a775: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::fc20:fbff:fef0:79e4  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether fe:20:fb:f0:79:e4  txqueuelen 0  (Ethernet)\n        RX packets 133499  bytes 23048506 (23.0 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 139255  bytes 60785977 (60.7 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethea6abd4d: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::245a:9ff:fee5:49cf  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 26:5a:09:e5:49:cf  txqueuelen 0  (Ethernet)\n        RX packets 300496  bytes 88641425 (88.6 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 280979  bytes 143269709 (143.2 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethef5d1630: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::848e:c5ff:feec:ab0b  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 86:8e:c5:ec:ab:0b  txqueuelen 0  (Ethernet)\n        RX packets 40794  bytes 7221671 (7.2 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 40380  bytes 14718234 (14.7 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethf1e8b5ff: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::7c0d:62ff:fecc:2b93  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 7e:0d:62:cc:2b:93  txqueuelen 0  (Ethernet)\n        RX packets 74813  bytes 7327271 (7.3 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 84111  bytes 7426273 (7.4 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethf8c3c8b0: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::1c11:99ff:fe3a:b803  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 1e:11:99:3a:b8:03  txqueuelen 0  (Ethernet)\n        RX packets 393730  bytes 67763877 (67.7 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 400020  bytes 87045635 (87.0 MB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\nvethfd5fb0b1: flags=4163\u003cUP,BROADCAST,RUNNING,MULTICAST\u003e  mtu 1450\n        inet6 fe80::18e4:14ff:fe28:2be3  prefixlen 64  scopeid 0x20\u003clink\u003e\n        ether 1a:e4:14:28:2b:e3  txqueuelen 0  (Ethernet)\n        RX packets 128519  bytes 302536179 (302.5 MB)\n        RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 203565  bytes 1026685978 (1.0 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\n RX errors 0  dropped 0  overruns 0  frame 0\n        TX packets 203565  bytes 1026685978 (1.0 GB)\n        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n",
	}

	e := cloudevents.NewEvent()
	e.SetID("id")
	e.SetType("type")
	e.SetSource("source")
	if err := e.SetData(cloudevents.ApplicationJSON, data); err != nil {
		t.Fatal(err)
	}

	// Act
	event, err := Handle(context.Background(), e)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if event == nil {
		t.Errorf("received nil event") // fail on nil
	}
	fmt.Println(event)
}