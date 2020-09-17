FROM hybridgroup/gocv AS dev
ENV GOPROXY=https://goproxy.cn
WORKDIR /app
ADD . /app
RUN cd /app && go build -o goapp
CMD [ "goapp" ]

# FROM alpine:latest
# RUN apk update && \
#    apk add ca-certificates && \
#    update-ca-certificates && \
#    apk add --no-cache beanstalkd && \
#    rm -rf /var/cache/apk/*

# COPY --from=development /app/goapp /app
# WORKDIR /app
# EXPOSE 11300


