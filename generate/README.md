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


# Test with postgres

Modify postgres.conf as following:

```
ssl = on
ssl_ca_file = '/home/gpadmin/workspace/zhaorui_self_practice/generate/root.ca'
ssl_cert_file = '/home/gpadmin/workspace/zhaorui_self_practice/generate/server.crt'
#ssl_crl_file = ''
ssl_key_file = '/home/gpadmin/workspace/zhaorui_self_practice/generate/server.key'
```

Modify pg_hba.conf

```
hostssl all     testssl    0.0.0.0/0       md5     clientcert=verify-ca|verify-full
```
See detailed settings at [auth-options](https://www.postgresql.org/docs/current/auth-pg-hba-conf.html#:~:text=20.14%20for%20details.-,auth%2Doptions,-After%20the%20auth)

Create user:
```
CREATE USER testssl WITH PASSWORD '123456' LOGIN; 
```

Restart postgres and then connect with psql:

```
export PGSSLROOTCERT=$(readlink -f root.ca)
export PGSSLCERT=$(readlink -f client.crt)
export PGSSLKEY=$(readlink -f client.key);
psql "sslmode=verify-full  dbname=postgres" -U testssl -h centos7gpdb7
```
See detailed settings at [sslmode](https://www.postgresql.org/docs/8.4/libpq-connect.html#LIBPQ-CONNECT-SSLMODE:~:text=server%20debug%20output.-,sslmode,-This%20option%20determines)
