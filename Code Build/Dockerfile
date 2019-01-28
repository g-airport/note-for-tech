#---------------------------------------------------*
# chapter 1
#---------------------------------------------------*
FROM debian:latest
WORKDIR /application
RUN echo "Asia/Shanghai" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && date
COPY ./dist/server/application/bin/${PROJECT_NAME} /application/bin/
WORKDIR /application/bin

ENTRYPOINT ["./project-name"]

#---------------------------------------------------*
# chapter 2
#---------------------------------------------------*
FROM docker-image-repo as builder

FROM docker-image-repo/lib/golang-runtime:debian

## Args
#ARG ARG_CI_BUILD_INFO
ARG ARG_PROJECT_GROUP=project-group
ARG ARG_PROJECT_NAME


ENV PROJECT_GROUP $ARG_PROJECT_GROUP
ENV PROJECT_NAME $ARG_PROJECT_NAME
ENV SERVER_ADDR=:0
#Gitlab-ci
#ENV CI_BUILD_INFO $ARG_CI_BUILD_INFO

# project-code-repo
COPY --from=builder /go/src/project-repo/${PROJECT_GROUP}/${PROJECT_NAME}/bin/${PROJECT_NAME} /application/bin/