# go-webserver
A simple webserver with HTTPS support written in go

To compile the binary, run `go build .`.

## Configuration
The configuration is done with the config.yaml file, which must be in the same location as the binary. Example config file:

```
  - file: "index.html"
    ipaddress: "0.0.0.0"
    port: 8080
    tls: true
    cert: "localhost.pem"
    key: "localhost-key.pem"
  - file: "index1.html"
    ipaddress: "127.0.0.1"
    port: 8081
    tls: false
    cert: ""
    key: ""
```
