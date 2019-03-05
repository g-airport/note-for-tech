## 

server module & event driver & process communicate

- nginx server module
- web request processing mechanism
- process function && communicate

### structure

- core
- event 
- http
- mail
- misc # other module
- os

### std . http
- ngx_http_core # config port , url analyze , server response err handle 
- ngx_http_auth_basic_module # base http auth
- ngx_http_access_module # base ip access control strategy
- ngx_http_autoindex_module # handle "/" req and auto generate category list
- ngx_http_browser_module # decode value of http header "User-Agent"
- ngx_http_charset_module #  index code  

### web request mechanism

- multi-process
        every time server accept one client req -> server main process -> fork child process 
        client  -> when user closed 
        
        fast-speed dependency each child process ,
        but access rate bigger, source will dead && can't support new req

- multi-thread 
        source less, source share ,when process down ,it's hard to recover soon
        
### work process

- main

    read nginx conf && verify valid and correct \
    establish, bind, close socket conns \
    dep conf generate , manege, stop work process
    
    accept external instruction eg: restart , update service \
    even failed , roll back  
    
    not interrupt make smooth update && restart new conf \
    open access log , acquire file description symbol
    
    compile && handle perl script
    
                        <fork>

- worker
    
    accept client request handle that function module \
    io call , acquire response data \
    communicate with backend server && accept res 
    
    cache data && access data index && select ,call that
    
    send req response , ack client req 
    
    accept instruction from main eg: restart , update , exit
    
### compress && cache

- ngx_http_gzip_module

```bash
gzip on | off
gzip_buffers number | size # default number * size 128
gzip_comp_level 1 ~ 9 # default 4 or 5 , high -> cpu
gzip_http_verson 1.0 | 1.1
gzip_min_length 1024
gzip_proxied off | any # all data from backend
gzip_types text/plain application/x-javascript text/css application/xml
gzip_vary on | off # vary:Accept-Encoding
```

- ngx_http_gzip_static_module
- ngx_http_gunzip_module


    
        
        