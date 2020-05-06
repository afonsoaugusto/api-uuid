# api-uuid

Example api-uuid with Elastic_APM_Server plugin

```sh
docker build -t api-uuid .
docker run --rm -it -p 8080:8080 \
-e ELASTIC_APM_SERVICE_NAME=api-uuid \
-e ELASTIC_APM_SERVER_URL=http://apm-server:8200 \
api-uuid
```

```sh
> $ curl -i http://localhost:8080/
HTTP/1.1 200 OK
Date: Wed, 06 May 2020 20:15:24 GMT
Content-Length: 36
Content-Type: text/plain; charset=utf-8

6252bbd4-b77f-435b-b036-a8acbfd208cf%  

> $ curl -i http://localhost:8080/health
HTTP/1.1 200 OK
Date: Wed, 06 May 2020 20:15:45 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OK%
```
