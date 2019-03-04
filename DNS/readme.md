
### DNS RR (Resource Records)
---
- **owner** 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Indicate the DNS domain name that owns the resource record
- **TTL**  

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; This field is optional for most resource records. Indicate how long other DNS servers cache them before expiring the record. Resource records with a TTL value of zero will not be cached
- **CLASS** 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; IN（Internet most situation ）CS（CSNET）、CH（CHAOS）、HS（Hesiod）
- **TYPE** 

| Type | Usage | Example |
| --- | --- | --- | 
|A       | Host Address | owner  TTL  CLASS  A  IPv4_address |
|AAAA    | IPv6 Host Address | owner  TTL  CLASS  AAAA  IPv6_address |
|NS      | Authoritative name server | owner  TTL  CLASS  NS name_server_domain_name |
|MD      | Mail destination(Deprecated，use MX) | |
|MF      | Mail forwarder(Deprecated，use MX) | |
|CNAME   | Regular name of the alias | owner  TTL  CLASS  CNAME  canonical_name |
|SOA     | Mark the beginning of the authoritative area | owner  TTL  CLASS SOA  name_server  responsible_person(serial_number  refresh_interval  retry_interval  expiration  minimum_time_to_live) |
|MB      | Mail domain (beta) | |
|MG      | Mail group member(beta) | |
|MR      | Mail rename domain(beta) | |
|NULL    | NULL RR(beta) | |
|WKS     | Well-known business description | |
|PTR     | Domain pointer | owner  TTL  CLASS  PTR targeted_domain_name |
|HINFO   | Host info | |
|MINFO   | Mailbox or mailing list information |
|MX      | Mail Exchange | owner  TTL  CLASS  MX  preference mail_exchanger_host |
|TXT     | Text String |
- **RDATA**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; The necessary fields for describing the information of the resource and having a variable length, which vary with the CLASS and TYPE.

---

A: Host Addr (A) source record  DNS domain map Internet protocol (IP) version 4 32 bit address（RFC 1035）

AAAA: IPv6 Host Addr (AAAA) source record DNS map  Internet protocol (IP) version 6 128 bit address（RFC 1886）

NS: owner mark -> DNS map name_server_domain_name Segment HostName of the running DNS server

CNAME: routine (CNAME) source record 。owner segment of alias or backup DNS domain map canonical_name Segment Std or Main DNS domain , the data use std or main dns necessarily and valid dns domain 

SOA: Starting authority(SOA) , point to domain source name including source of info server name and representing the place base attribute ;SOA recording any std place of initial record; usage of save area update or expiration or time ticking  and so on

PTR : PTR like targeted_domain_name explained from owner's DNS namespace's another place , eg: in-addr.arpa special usage ，To provide a reverse lookup of the address-name mapping. In most cases, each record provides information that points to another DNS domain name, such as the corresponding host (A) address resource record in the forward lookup zone (RFC 1035)

MX: Mail Exchange (MX) mail_exchanger_host provided for mail exchanger that supporting mail router in order to send email to owner segment of domain ; preference represent order of sending. Every Exchange Host have the adaptor of Host Address Record （RFC 1035）

---

### SOA Record

| name | usage |
| --- | --- |
| name_server           | The primary name server for the zone |
| responsible_person    | The email address of the person who manages the area (the first one "." If you change to "@") |
| serial_number         | Serial number, all data applied to the area, usually used time YYYYMMDDHHmm, for primary and secondary synchronization, except for the very early bind version (bind4.8.3), can use m for minutes, h for hours, d for days, w for week |
| refresh_interval      | How often does the secondary server check if the data in the area is up to date  |
| retry_interval        | Unable to connect to the primary server after the refresh time, how often try to reconnect |
| expiration            | After the expiration, the main name server cannot be connected within the expiration time, and the secondary name server invalidates the area. |
| minimum_time_to_live  | Bind8.2 indicates the minimum default TTL value and cached negative TTL before bind. After bind8.2, minimum_time_to_live indicates that the buffer is negative TTL. |





###ROOT DNS Server

| Server | IP option | Address |
| --- | --- | --- |
| A.root-servers.net  |   A      |    198.41.0.4 United States（support IPv6）|
| A.root-servers.net  |   AAAA   |    2001:503:ba3e::2:30|
| B.root-servers.net  |   A      |    192.228.79.201 United States（support IPv6）|
| B.root-servers.net  |   AAAA   |    2001:500:84::b|
| C.root-servers.net  |   A      |    192.33.4.12 法国（support IPv6）|
| C.root-servers.net  |   AAAA   |    2001:500:2::c|
| D.root-servers.net  |   A      |    199.7.91.13 United States（support IPv6）|
| D.root-servers.net  |   AAAA   |    2001:500:2d::d|
| E.root-servers.net  |   A      |    192.203.230.10 United States|
| F.root-servers.net  |   A      |    192.5.5.241United States（support IPv6）|
| F.root-servers.net  |   AAAA   |    2001:500:2f::f|
| G.root-servers.net  |   A      |    192.112.36.4 United States|
| H.root-servers.net  |   A      |    128.63.2.53 United States（support IPv6）|
| H.root-servers.net  |   AAAA   |    2001:500:1::803f:235|
| I.root-servers.net  |   A      |    192.36.148.17 Sweden（support IPv6）|
| I.root-servers.net  |   AAAA   |    2001:7fe::53|
| J.root-servers.net  |   A      |    192.58.128.30 United States（support IPv6）|
| J.root-servers.net  |   AAAA   |    2001:503:c27::2:30|
| K.root-servers.net  |   A      |    193.0.14.129 United Kingdom（support IPv6）|
| K.root-servers.net  |   AAAA   |    2001:7fd::1|
| L.root-servers.net  |   A      |    199.7.83.42 United States（support IPv6）|
| L.root-servers.net  |   AAAA   |    2001:500:3::42|
| M.root-servers.net  |   A      |    202.12.27.33 Japan（support IPv6）|
| M.root-servers.net  |   AAAA   |    2001:dc3::35|



