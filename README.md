# HTTP Test

## Status codes
`GET` `/status/{StatusCode}` Return status code

`POST` `/status/{StatusCode}` Return status code

## Random 502 （Internal Server Error）
To test outlier detection, which is supported by istio, `/rand_status` will randomly respond to `200` or `500` to simulate that encountered the error. Using it before set error rate, if `ERR_RATE` not set, the default is 50 (i.e. 50%).
 
SET `env` `ERR_RATE` [0, 100]

`GET` `/rand_status`

