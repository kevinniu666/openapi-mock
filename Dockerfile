FROM golang:1.16.9 AS build
COPY . /src
WORKDIR /src
ENV GOPROXY=https://goproxy.cn
RUN go mod vendor && go build -o app main.go

FROM centos:7.9.2009
COPY --from=build /src/app /app
RUN chmod +x /app
CMD /app
EXPOSE 8888