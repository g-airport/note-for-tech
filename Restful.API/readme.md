## Restful API
(Representational State Transfer)

### RFC 3986 
定义了通用的URI语法

```text
URI = scheme “://” authority “/” path [ “?” query ][ “#” fragment ]
```

- scheme: 指底层用的协议，如http、https、ftp \
- host: 服务器的IP地址或者域名 \
- port: 端口，http中默认80 \
- path: 访问资源的路径，就是咱们各种web 框架中定义的route路由 \
- query: 为发送给服务器的参数 \
- fragment: 锚点，定位到页面的资源，锚点为资源id

### Method
- GET：读取（Read）
- POST：新建（Create）
- PUT：更新（Update）
- PATCH：更新（Update），通常是部分更新，向客户端提供改变的属性
- DELETE：删除（Delete）

**属性覆盖：**\
> X-HTTP-Method-Override

```bash
POST /api/Person/4 HTTP/1.1  
X-HTTP-Method-Override: PUT
```

```bash
GET: 200 OK
POST: 201 Created
PUT: 200 OK
PATCH: 200 OK
DELETE: 204 No Content
```

**400 Bad Request：** 服务器不理解客户端的请求，未做任何处理 \
**401 Unauthorized：** 用户未提供身份验证凭据，或者没有通过身份验证 \
**403 Forbidden：** 用户通过了身份验证，但是不具有访问资源所需的权限 \
**404 Not Found：** 所请求的资源不存在，或不可用 \
**405 Method Not Allowed：** 用户已经通过身份验证，但是所用的 HTTP 方法不在他的权限之内 \
**410 Gone：** 所请求的资源已从这个地址转移，不再可用 \
**415 Unsupported Media Type：** 客户端要求的返回格式不支持。比如，API 只能返回 JSON 格式，但是客户端要求返回 

### Summary

（1）每一个URI代表一种资源；

（2）客户端和服务器之间，传递这种资源的某种表现层；

（3）客户端通过四个HTTP动词，对服务器端资源进行操作，实现"表现层状态转化"。