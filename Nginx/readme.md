## About Nginx

###Base

#### Intro

- www.nginx.org

- single thread , single process

- cpu bound like mysql

- as http server 

- reverse proxy 

- FastCGI/SSL/Virtual Host/URL Rwrite/Gzip/HTTP Basic Auth

#### Guide 

- install 

```bash
# acquire 
# gcc GNU Compiler Collection 
# automake Makefile.am -> Makefile.in dep perl (libtool)
# pcre-devel PCRE Perl Compatible Regular Expression (Rewrite HTTP core dep PCRE)
# zlip (zlib-devel) Compress Function
# openssl open-ssl-devel 
wget http://nginx.org/download/nigix-x.x.x.tar.gz
tar zxvf nginx-x.x.x.tar.gz
cd nginx-x.x.x
ls && make
```
- configure

```bash
./configure \
--prefix=/usr/local/nginx \
--sbin-path=/usr/local/nginx/sbin/nginx \
--conf-path=/usr/local/nginx/conf/nginx.conf \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--pid-path=/var/run/nginx/nginx.pid \
--lock-path=/var/lock/nginx.lock \
--user=nginx \
--group=nginx \
--with-http_ssl_module \
--with-http_stub_status_module \
--with-http_gzip_static_module \
--http-client-body-temp-path=/var/tmp/nginx/client/ \
--http-proxy-temp-path=/var/tmp/nginx/proxy/ \
--http-fastcgi-temp-path=/var/tmp/nginx/fcgi/ \
--http-uwsgi-temp-path=/var/tmp/nginx/uwsgi \
--http-scgi-temp-path=/var/tmp/nginx/scgi \
--with-pcre

make install
```

-catalog 


- conf  #core configuration file \
- html  #nginx web file 50x err notify index \
- logs \
- sbin upstart script file ,accept different param to make different feature  


- debug

```bash
cat /var/run/nginx/nginx.pid # random pid every start

kill -QUIT pid  # 平缓关闭，不再接受新的请求，处理完当前请求后关闭
kill -TERM pid  # 快速停止 
kill -HUP pid   # 使用新配置文件启动，平缓停止原有进程，平滑重启
kill -USR1 pid  # 重新打开配置文件，用于 日志切割
kill -USR2 pid  # 使用新版本配置文件启动，然后平缓停止原有进程，平滑升级
kill -WINCH pid # 平滑停止工作进程，用于平滑升级
```

- rotate log

```bash
#!/bin/bash
PID=`cat /var/run/nginx/nginx.pid`
mv /var/log/nginx/access.log /var/log/nginx/`date +%Y_%m_%d:$H:$M:$S`.access.log
kill -USR $PID
``` 

- module

```bash
grep -v "#" nginx.conf | grep -v "^$"

#script nginx.conf
work_process 1; # default start one process 

event {
# events effect nginx server and user net connection 
# eg: many connects at the same time
# use which model to deal request 
# every work process fit how many connection 
# sequence net connect or not
}

http {
# http core cache,proxy,format log
# http including many #server model that including many location model
# configure file import , MIME-Type definition , log def , is send file or not
# connect timeout, single connection req-limit

include mime.types;
default_type application/octet-stream; # default file type
sendfile on; # zero copy --> optimism  read nginx  if download --> off 
sendfile_max_chunk 512k; # file length < 512k http/server/location
worker_rlimit_nofile 65535; # linux max open file number
keepalive_timeout 65 ; # unit s default 1.8.1 < 120s
    Keep-Alive:timeout=60 # browser recv context from server
    Connection:close # above 

server {
    # set vm host,
    listen 8090; # server global port 
    # listen ip:port;
    # Unix:/www/file # unix socket
    server_name localhost; # when access this server name call this internal config
    #server_name local1 local2 local3
    #server_name x.x.x.x # localhost net ip
    location / #regular  {
            # a command of server 
            # request string UIL deal including redirect ,data cache, ack
            root html; # default index path (opt: absolute path)
            index index.html index.html;
        }
        
    # =   #before std uri need req string match uri absolutely -> deal request 
    # ~   #identify utter or lower
    # ~*  #no identify utter
    # !~  #identify utter or lower not match
    # !~* #!identify utter or lower not match
    # ^   #match start with
    # $   #match end with
    # \   #escape character . * ? ...
    # *   #represent infinite character && length
    # 
    # -f and !-f #exist file
    # -d and !-d #exist path
    # -e and !-e #exist path and file
    # -x and !-x #execute or not
    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        # location deal different err code 
        root html; # def path of 50x
        }    
    }
}
```

```bash
user xxxx; # every instruction use symbol ";" end
work_process 1; # auto # current cpu number
work_cpu_affinity 0001 0010 0100 1000; # 4 core
# 00000001 00000010 00000100 00001000 00010000 00100000 01000000 10000000；# 8 core
# find cpu number
grep process /proc/cpuinfo | wc -l
```

```bash
# pid logs/nginx.pid # (option absolute path)
# error_log logs/error.log
# error_log logs/error.log notice ; # log level
# error_log logs/error.log info ;
```

- grammar 

```bash
error_log file [debug | info | notice | warn | error | crit] | \ 
[{debug_core | debug_alloc | debug_mutex | debug_event | debug_http | debug_mail | debug_mysql}]
```

- include file;
```bash
inlcude /usr/local/nginx/conf.d/some.conf
grep -v "#" conf.d/some.conf | grep -v "^$"
```

### Nginx Optimism 

- net work connect

```bash
events {
    
    accept_mutex on; # only one request awake many sleep process at the same time
    multi_accept on; # accept many conns req function 
    
    use epoll; # use event driver, high performance
    
    worker_connections 1024 #set unit work process max conns 1024  
}

http {
    server_tokens off; # unknown CVE , aviod hiking 
}

# definition kind of source , third party
include mime.types; # net source media type HTML/GIF/XML/FLASH
default_type application/octet-stream;

# access log concrete info (log level) 
# one error_log but many different server

log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                '$status $body_bytes_sent "$http_referer" '
                '"$http_user_agent" "$http_x_forwarded_for"';
access_log /var/log/nginx/access.log main;
```

- json log

```bash
log_format logstash_json '{"@timestamp":"$time_iso8601",'
        '"host":"$server_addr",'
        '"clientip":"$remote_addr",'
        '"size":$body_bytes_sent,'
        '"responsetime":$request_time,'
        '"upstreamtime":"$upstream_response_time",'
        '"upstreamhost":"$upstream_addr",'
        '"http_host":"$host",'
        '"url":"$uri",'
        '"domain":"$host",'
        '"xff":"$http_x_forwarded_for",'
        '"referer":"$http_referer",'
        '"agent":"$http_user_agent",'
        '"status":"$status"}';
server {
　　　　listen 8090;
　　　　server_name some_server;
　　　　access_log /var/log/nginx/some_server.access.log logstash_json;
　　　　location / {
　　　　　　root html;
　　　　　　index index1.html index.htm;
　　　　}
　　　　error_page 500 502 503 504 /50x.html;
　　　　location = /50x.html {
　　　　root html;
　　　　　}
　　}
```

- http code

```bash
200 #success
301 #permantly redirect 
302 #templately redirect
403 #forbidden access，auth deny
400 #err request grammar problem  
403 #auth can't work
404 #not found
500 #server internal error
501 #did not set the website being visited as the content requested by the browser
502 #net gateway
503 #none available 
504 #gateway time out , not complete 
# the processing request within the specified time or server overload 
505 #not support http (ＨＴＴＰ/1.1)
```

- sysctl.conf

```bash
sysctl -a | grep max_backlog
# net interface deal speed faster than kernel deal, send queue number
net.core.netdev_max_backlog = 1000 # default , make some big value
# system adapt tcp conns case 1: conns timeout case 2 : repeated send
net.core.somaxconn = 128 # high currency use big value 🤣
# set top socket just not match any file handle, if over -> reset tcp_max_orphans
# avoid DDOS when memory satisfied . can set bigger
net.ipv4.tcp_max_orphans = 32768;
# record un ack client conns req number
net.ipv4.tcp_max_syn_backlog = 256 # set big better
# timestamp , avoid sequence number overlapped 
# default none sequence bag
net.ipv4.tcp_timestamp = 1 # optimism 0
# kernel abandon before tcp connect syn+ack bag number allow 1 means one connect
net.ipv4.tcp_synack_retries=5 # use 1 , avoid syn attack
net.ipv4.tcp_syn_retries = 5 
```

- cpu optimism

```bash
work_process 1; # auto # current cpu number
work_cpu_affinity 0001 0010 0100 1000; # 4 core
```

- net config
```bash
keepalived_timeout 60 50; # nginx to client , Keep-Alive msg Header && browser to server
sendtime 10s # http core instruction timeout not enter established status, just two hand shake
# if no ack ,nginx will closee client conn
client_header_timeout # head_buffer length 1kb enough if div or cookie 4K
multi_accept on; # 

```

- command
```bash
use 
work_process
work_connections
work_rlimt_sigpending 65535; # every event process queue length. use poll;
# devpoll_changes server to kernel pass event number
# devpoll_changes server from kernel read event number 
devpoll_changes && devpoll_changes # config /dev/poll default 512
kqueue_changes && kqueue_events # default 512
epoll_events # 512
rtsig_signo # rtsig mode 
rtsig_overflow #rtsig_signo  
# config Nginx rtsig mode 
# the first of the two signals used, 
# the second signal is incremented by one on the number of the first signal.
 
rtsig_overflow #rtsig queue over 
# When an overflow occurs when nginx flushes the rtsig queue, 
# they will continuously call poll() and rtsig.poll() 
# to handle outstanding events.

# Until rtsig is drained to prevent new overflows, 
# when the overflow is processed, nginx enables rtsig mode again.

# rtsig_overflow_events specifies poll() event number ，default 16，

# rtsig_overflow_test 
# Specify how many events poll() handles, 
# and nginx will empty the rtsig queue. The default value is 32.

# rtsig_overflow_threshold Can only run under the Linux 2.4.x kernel.
# Before emptying the rtsig queue, nginx checks the kernel to 
# determine how the queue is filled. The default is 1/10.
# “rtsig_overflow_threshold 3” mean 1/3

```