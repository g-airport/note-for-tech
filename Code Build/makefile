# 基于GOLANG项目构建
PROJECT_NAME = $SERVER_NAME
MAIN_PATH = ./src/main.go
SERVER_BIN_DIST_DIR = ./dist/server/application/bin
CURRENT_VERSION := `head -1 release_log `
BRANCH = $(shell basename `git symbolic-ref HEAD`)
COMMIT = $(shell git rev-parse HEAD)

build:clean
	GOOS=linux GOARCH=amd64 go build -o $(SERVER_BIN_DIST_DIR)/${PROJECT_NAME} ${MAIN_PATH}
	# go build -o $(SERVER_BIN_DIST_DIR)/${PROJECT_NAME} ${MAIN_PATH}
	cd $(SERVER_BIN_DIST_DIR) && chmod +x ${PROJECT_NAME}

clean:
	rm -rf dist/server

# 仅编译镜像
build-image:
	docker build --build-arg build=build --rm -t $(PROJECT_NAME):$(CURRENT_VERSION) --build-arg ARG_PROJECT_NAME=${PROJECT_NAME} ./

# 项目发布
publish:build
	docker build --rm -t $(PROJECT_NAME):$(CURRENT_VERSION) --build-arg ARG_PROJECT_NAME=${PROJECT_NAME} ./
	docker tag $(PROJECT_NAME):$(CURRENT_VERSION) $(REPO_URL)$(PROJECT_NAME):$(CURRENT_VERSION)
	docker login -u $(REPO_USER) -p $(REPO_PASS) $(REPO_URL) && docker push $(REPO_URL)$(PROJECT_NAME):$(CURRENT_VERSION)

# Docker Repo Info
auth-info
    echo $(REPO_USER)
	echo $(REPO_PASS)
	echo $(REPO_URL)

.PHONY: gen build build-image publish clean auth-info