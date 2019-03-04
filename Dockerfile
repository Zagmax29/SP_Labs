FROM alpine
RUN apk add --no-cache git make musl-dev go
RUN apk update go
WORKDIR /home
COPY task_17.go /home/task_17.go
ENTRYPOINT ["go"] 

