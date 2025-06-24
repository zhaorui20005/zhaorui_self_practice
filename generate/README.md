# Generate the root.ca server.key server.crt client.key client.crt

```
go build ./generate_cert.go

./generate.go

```
# Test with openssl tools

Start server side:
```
openssl s_server -key server.key  -cert server.crt -CAfile root.ca  -accept 4433
```

Start client side:

```
openssl s_client -connect 127.0.0.1:4433 -cert client.crt -key client.key -CAfile root.ca
```