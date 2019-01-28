## Chapter 1 : makefile


#### 项目一般结构


- ./src 项目源码
- ./bin go build 二进制文件
- ./log 项目日志

---

```bash
GO = go
GO_GET = $(GO) get
BUILD = ${CROSS} ${GO} build ${GO_FLAG}

ENV_PATH = ./
VERSION_NAMESPACE = ${ENV_PATH}
BUILD_VERSION = ${BUILD} -ldflags "-X $(VERSION_NAMESPACE).repository=$(REPO) -X $(VERSION_NAMESPACE).branch=$(BRANCH) -X $(VERSION_NAMESPACE).commit=$(COMMIT)"

```

**Docker镜像仓库信息**

>支持export

- REPO_USER=xxx
- REPO_PASS=xxx
- REPO_URL=xxx

## Chapter 2 : Dockerfile
