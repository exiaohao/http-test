# HTTP Test

## Status codes
`GET` `/status/{StatusCode}` Return status code

`POST` `/status/{StatusCode}` Return status code

## Random 502 （Internal Server Error）
Test outlier detection, which is supported by istio. Using it before set error rate, if `ERR_RATE` not set, the default is 50 (i.e. 50%).

`env` `ERR_RATE` 0-100

`GET` `/rand_status`

