# Demo Load Balance with HAProxy and Golang

With this demo, I used Golang to create 3 webservers. HAProxy is used to create reverse proxy.

3 webservers will be started at port **:9000**, **:9001**, **:9002**

HAProxy will be started at port **:8012**

## What you can learn?
1. How to create multiple web servers on 1 file main.go
2. How to config HAProxy to be reverse proxy

## Use
Run: `go run main.go` to start 3 webservers

Run: `sudo haproxy -f haproxy.cfg`

Run: `sudo brew services restart haproxy`

Enter url: `localhost:8012` on browser, then refresh several times to view the result

View statistic at: `localhost:8012/stats`

## My article about this source code:
Vietnamese: https://bit.ly/37DWzHN

English: comming soon