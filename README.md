# srvlb
Dns SRV loadbalancer written in Go

Tested for integration with consule

Example
```
docker run -d -e CONSUL_ALLOW_PRIVILEGED_PORTS= --net=host --name=consul_server consul:v0.7.0 agent -ui -client=172.17.0.1 -server -dns-port=53 -bind=172.17.0.1 -bootstrap
docker build -t srvlb .
docker run -d -p 172.17.0.1:1080:1080 --name srvlb --dns 172.17.0.1 --dns-search service.consul srvlb
docker run -dt --name http_2 -p 172.17.0.1:8001:8000 python:3.5-alpine python -m http.server
docker run -dt --name http_1 -p 172.17.0.1:8002:8000 python:3.5-alpine python -m http.server
curl -vd '{"ID": "http_1", "Name": "http", "Address": "172.17.0.1", "Port": 8001}' 172.17.0.1:8500/v1/agent/service/register
curl -vd '{"ID": "http_2", "Name": "http", "Address": "172.17.0.1", "Port": 8002}' 172.17.0.1:8500/v1/agent/service/register
http_proxy=http://172.17.0.1:1080 curl http
http_proxy=http://172.17.0.1:1080 curl http
docker logs http_1
docker logs http_2

```
