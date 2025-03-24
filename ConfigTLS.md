To create a self signed certificate for client and server, you may need to config by the following scripts.

## client_config
```
[req]
distinguished_name = req_distinguished_name
req_extensions = req_ext
prompt = no
 
[req_distinguished_name]
CN = rocky9
 
[req_ext]
subjectAltName = IP:127.0.0.1,DNS:client_hostname
```

## server_config

```
[req]
distinguished_name = req_distinguished_name
req_extensions = req_ext
prompt = no
 
[req_distinguished_name]
CN = rocky9
 
[req_ext]
subjectAltName = IP:127.0.0.1,DNS:server_hostname
```

## script content
```
cat generate_certs
openssl genrsa -out root.key 2048
openssl req -x509 -new -nodes -key root.key -sha256 -days 36500 -out root.crt -subj "/C=CN/ST=Beijing/L=Beijing/O=vmware/OU=cibg/CN=vmware.com"
 
openssl genrsa -out server.key 2048
openssl rsa -in server.key -pubout -out server.pub
openssl req -new -key server.key -out server.csr -config server_config -extensions req_ext
openssl x509 -req -in server.csr -CA root.crt -CAkey root.key -CAcreateserial -out server.crt -days 36500 -sha256 -extfile server_config -extensions req_ext
 
openssl genrsa -out client.key 2048
openssl rsa -in client.key -pubout -out client.pub
openssl req -new -key client.key -out client.csr -config client_config -extensions req_ext
openssl x509 -req -in client.csr -CA root.crt -CAkey root.key -CAcreateserial -out client.crt -days 36500 -sha256 -extfile client_config -extensions req_ext
 
chmod 0600 server.crt root.crt
```
