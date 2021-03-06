
# Minio Usage

## run minio server on docker

```shell
docker run \
  --name minio -d \
  -p 9000:9000 \
  -p 9001:9001 \
  -v data:/data \
  -e "MINIO_ACCESS_KEY=pms" \
  -e "MINIO_SECRET_KEY=pms123456" \
  --restart=always \
  minio/minio server /data --console-address ":9001"
```


## run minio server on docker (host: ARMv7)

```shell
docker run --name minio -d \
        --restart=awlays \
        -p 9999:9000 \
        -e "MINIO_ACCESS_KEY=pms" \
        -e "MINIO_SECRET_KEY=pms123456" \
        -v data:/data 
        minio/minio:RELEASE.2020-12-03T05-49-24Z server /data
```
