## Build image
``` bash
docker build -t stresstest .
``` 

## Run command
```bash
docker run stresstest stress -u https://httpbin.org/get -r 3500 -c 200
``` 

## Details
```bash
-u https://httpbin.org/get  --URL
-r 3500                     --Requests
-c 200                      --concurrency
```