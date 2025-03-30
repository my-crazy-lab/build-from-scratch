## Read later
- [The big-endians and the little-endians](https://www.rfc-editor.org/ien/ien137.txt)
- [BNF](https://datatracker.ietf.org/doc/html/rfc5234.html)
- [The Domain Naming Convention for Internet User Applications](https://datatracker.ietf.org/doc/html/rfc819.html)
- [DNS performance](https://ieeexplore.ieee.org/document/1041066)
- [DNS security](https://datatracker.ietf.org/doc/html/rfc4033.html)
- [Internet mail architecture](https://datatracker.ietf.org/doc/html/rfc5598.html)
- [Internet message format](https://datatracker.ietf.org/doc/html/rfc5322.html)
- [MIME 1](https://datatracker.ietf.org/doc/html/rfc2045.html)
- [MIME 2](https://datatracker.ietf.org/doc/html/rfc2046.html)
- [base16, base32, base64](https://datatracker.ietf.org/doc/html/rfc4648.html)
- [SMTP](https://datatracker.ietf.org/doc/html/rfc821.html)
- [IMAP](https://datatracker.ietf.org/doc/html/rfc3501.html)
- [POP](https://datatracker.ietf.org/doc/html/rfc1939.html)
- [HTTP](https://datatracker.ietf.org/doc/html/rfc2616.html)
- [FTP](https://datatracker.ietf.org/doc/html/rfc959.html)
- [SSH](https://datatracker.ietf.org/doc/html/rfc4251.html)
- [URI](https://datatracker.ietf.org/doc/html/rfc3986.html)
- [Cookie](https://datatracker.ietf.org/doc/html/rfc6265.html)
- [UDP](https://datatracker.ietf.org/doc/html/rfc768.html)
- [RPC](https://datatracker.ietf.org/doc/html/rfc5531.html)
- [NFS](https://datatracker.ietf.org/doc/html/rfc1094.html)

## Introduction 
- LAN 
- MAN 
- WAN 
- Full mesh 
- Bus 
- Star
- Ring 
- Tree
- broadcast 
- unicast: one sender to on receiver 
- multicast: send to N receivers 
- anycast 

## Services and protocols 
- a user interacts with a service through a SERVICE ACCESS POINT 
- interaction between user and service by using primitives
- connectionless service: UDP 
  - reliable connectionless service: TCP
  - unreliable connectionless service: UDP

## OSI
- physical layer
  - unit of information exchanges is bit, rate is bit/s
  - electrical cable
  - optical fiber
  - wireless
- data link layer
  - unit of information exchanges is frame
  - a frame is a finit sequence of bit
- network layer
  - exchanges packets
  - a packet is a finit sequence of bytes, transport by datalink layer inside frames
- transport layer
  - exchanges segments
  - a packet is a finit sequence of bytes, transport by network layer inside packets
- application layer
  - exchanges application data unit

## Application layer
- Big-Endians & Little-Endians
- ASCII
- BNF
  - use to define programming language (can define messages exchanges between L7)
- Electronic mail
  - SMTP
  - POP
  - IMAP