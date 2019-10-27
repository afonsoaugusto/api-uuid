# api-uuid

```sh
docker build -t api-uuid .
docker run --rm -it -p 8080:8080 \
-e ELASTIC_APM_SERVICE_NAME=api-uuid \
-e ELASTIC_APM_SERVER_URL=http://apm-server:8200 \
api-uuid
```
