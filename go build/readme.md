### go 构建 

#### 二进制文件

```
go build -o (catagory + binary file name) ./main.go

```

#### 构建服务[micro](https://github.com/micro/micro)服务 注册到[consul](https://github.com/hashicorp/consul)

`CONSUL_HTTP_ADDR(注册Consul地址)`
`registry_address(注册目标机器)`

```
CONSUL_HTTP_ADDR= address:port ./(catagory + binary file name)  --registry=consul --registry_address=address:port  --server=grpc --client=grpc

```

#### 本地构建Consul

```
consul agent -dev -server -client 0.0.0.0

```

#### Docker 步骤
1. ```mkdir docker```
2. ```touch Dockerfile```
3. ```touch docker-entrypoint.sh```
4. ```chmod +x docker-entrypoint.sh```

`Dockerfile`

```
FROM ip_address/library/golang-runtime:latest       //选择仓库地址 + golang运行runtime版本 

ARG ARG_PROJECT_NAME=project-srv                    //项目名称

ENV PROJECT_NAME=${ARG_PROJECT_NAME}
ENV CONSUL_HTTP_ADDR="consul_address"               //注册地址 运行时要注册的地址
ENV SERVER_ADDR=":0"

COPY docker/${PROJECT_NAME} /catagory               //项目构建目录
COPY docker/docker-entrypoint.sh /usr/local/bin/    //$GOROOT

ENTRYPOINT ["docker-entrypoint.sh"]

```


`docker-entrypoint.sh`                    

```
#!/bin/sh
set -e

exec "${PROJECT_NAME}" --registry=consul --registry_address=${CONSUL_HTTP_ADDR} --server=grpc --server_address=${SERVER_ADDR} --client=grpc --selector=cache
```
`Tip:`

```docker目录下：1.二进制文件 2.Dockerfile 3.docker-entrypoint.sh```


#### push Docker仓库
1  ```docker login [repo address]```                              //输入账号密码 (同github)

2. ```docker pull [repo address]/library/golang-runtime:latest``` //build docker image 前 , pull golang runtime version (对于第一次构建docker镜像)

3. ```docker build --build-arg ARG_PROJECT_NAME=[binary file] -t [repo address]/deploy/[binary file]:v1.0.0.0 ./``` //binary 是./docker 目录下的 二进制文件 \
ARG_PROJECT_NAME 是 Dockerfile 的 PROJECT_NAME

4 ```docker push [repo address]/eploy/[binary file]:v1.0.0.0```   //push 时可以指定tag

###### push 完成

`以上操作是把 项目 放到docker 仓库`
`具体运行步骤 要在 portainer.io 待续 、、、`
