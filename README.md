# HTTP Test

## How to use

image `reg.qiniu.com/hao/http-test:latest`

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: http-test-v1
  namespace: songhao
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: http-test
        version: v1
    spec:
      containers:
      - name: http-test
        image: reg.qiniu.com/hao/http-test:latest
        imagePullPolicy: Always
        env:
        - name: ERR_RATE
          value: "10"
        - name: GIN_MODE
          value: "release"
        - name: VERSION
          value: "v1"
        ports:
        - containerPort: 3000
```

## Version
To test different versions, set env `VERSION` and GET `/version` will show your version.

## Status codes
`GET` `/status/{StatusCode}` Return status code

`POST` `/status/{StatusCode}` Return status code

## Random 502 （Internal Server Error）
To test outlier detection, which is supported by istio, `/rand_status` will randomly respond to `200` or `500` to simulate that encountered the error. Using it before set error rate, if `ERR_RATE` not set, the default is 50 (i.e. 50%).
 
SET `env` `ERR_RATE` [0, 100]

`GET` `/rand_status`
`GET` `/api/demo`

## Call api
From ***http-test 1*** to ***http-test 2*** simulate that in-cluster calling.

SET `env` `TARGET_SERVICE` to define another http-test service, and call `/api/demo` for test.
