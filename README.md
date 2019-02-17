# HTTP Benchmarks

Simple benchmark results to compare different web technologies against real life scenarios.


## Contestants

Bellow are listed the tested technologies and some details about them:

- **golang-native**: native implementation of a web server in go. Go version: go1.11.5 linux/amd64.
- **golang-fasthttp**: web server implemented in go using the library [fasthttp](https://github.com/valyala/fasthttp). Fasthttp version: v1.2.0. Go version: go1.11.5 linux/amd64.
- **nginx**: one of the most widely used web servers for static content and proxying. Nginx version: nginx/1.14.0 (Ubuntu).
- **nodejs-native**: native implementation of a web server with nodejs. Node version: v10.15.1.
- **python-gunicorn-default**: python implementation being served by the gunicorn web server with it's default worker type (sync). Gunicorn version: 19.9.0. Python version: 3.7.2.
- **python-gunicorn-gevent**: python implementation being served by the gunicorn web server with the gevent worker type. Gunicorn version: 19.9.0. Gevent version: 1.4.0. Greenlet version: 0.4.15. Python version: 3.7.2.
- **python-uwsgi**: python implementation being served by the uwsgi web server. uWSGI version: 2.0.18. Python version: 3.7.2.
- **rust-hyper**: rust implementation using the library [hyper](https://github.com/hyperium/hyper). Hyper version: 0.12.24. Rust version: rustc 1.32.0 (9fda7c223 2019-01-16).


## Tests

Each implementation provided a set of endpoints to be consumed by a benchmarking tool, each doing a different thing:

- **/hello**: return 200 "hello" as fast as possible.
- **/file**: send `file.txt` as fast as possible.
- **/sleep/10ms**: sleep for 10ms and return 200 "ok".
- **/sleep/100ms**: sleep for 100ms and return 200 "ok".
- **/sleep/1s**: sleep for 1s and return 200 "ok".
- **/sleep/5s**: sleep for 5s and return 200 "ok".

The sleep requests are there to simulate some blocking work a real life endpoint might need to do, like a database query, read from a file, call an API, etc.


## Methodology

The servers were running in a $15 digital ocean droplet:

```
MEMORY    VCPUS     SSD DISK    TRANSFER    PRICE
2 GB      2 vCPUs   60 GB       3 TB        $15/mo
```

The client was running in a $40 digital ocean droplet:

```
MEMORY    VCPUS     SSD DISK    TRANSFER    PRICE
8 GB      4 vCPUs   160 GB      5 TB        $40/mo
```

The client is more powerful so it could overpower the server for sure. The tool used to benchmark was [wrk](https://github.com/wg/wrk). The same arguments were used to run all the tests:

```
wrk -t12 -c400 -d1m --timeout 10s <URL>
```

The servers were restarted between each test to ensure pending requests from the previous test wouldn't affect the next one.

The contestant **nginx** does not participate in the sleep tests.


## Results

The actual result outputs from `wrk` can be seen in the `results` directory. Bellow you'll find a summary of each test. Items are ordered from best to worst.

### /hello

```
golang-fasthttp            Requests/sec:  32435.66
golang-native              Requests/sec:  24840.89
rust-hyper                 Requests/sec:  21406.03
nodejs-native              Requests/sec:  18891.35
nginx                      Requests/sec:  15973.72
python-gunicorn-gevent     Requests/sec:   5887.06
python-gunicorn-default    Requests/sec:   5721.54
python-uwsgi               Requests/sec:   3198.35
```

### /file

```
rust-hyper                 Requests/sec:  23442.26
golang-fasthttp            Requests/sec:  18502.80
nginx                      Requests/sec:  17380.04
golang-native              Requests/sec:  14058.13
nodejs-native              Requests/sec:   8459.55
python-gunicorn-gevent     Requests/sec:   4152.58
python-gunicorn-default    Requests/sec:   2765.78
python-uwsgi               Requests/sec:   2547.30
```

### /sleep/10ms

```
nodejs-native              Requests/sec:  17542.24
golang-native              Requests/sec:  17069.23
golang-fasthttp            Requests/sec:  16028.32
python-gunicorn-gevent     Requests/sec:   6064.69
python-uwsgi               Requests/sec:    220.57
rust-hyper                 Requests/sec:    195.33
python-gunicorn-default    Requests/sec:    162.98
nginx                      Requests/sec:       N/A
```

### /sleep/100ms

```
nodejs-native              Requests/sec:   3886.39
golang-native              Requests/sec:   3864.66
golang-fasthttp            Requests/sec:   3853.86
python-gunicorn-gevent     Requests/sec:   3571.40
python-uwsgi               Requests/sec:     39.80
python-gunicorn-default    Requests/sec:     28.40
rust-hyper                 Requests/sec:     19.93
nginx                      Requests/sec:       N/A
```

### /sleep/1s

```
golang-fasthttp            Requests/sec:    389.63
golang-native              Requests/sec:    388.86
nodejs-native              Requests/sec:    388.76
python-gunicorn-gevent     Requests/sec:    380.80
python-uwsgi               Requests/sec:      3.99
python-gunicorn-default    Requests/sec:      2.95
rust-hyper                 Requests/sec:      2.00
nginx                      Requests/sec:       N/A
```

### /sleep/5s

```
golang-native              Requests/sec:     78.77
python-gunicorn-gevent     Requests/sec:     75.41
nodejs-native              Requests/sec:     73.56
golang-fasthttp            Requests/sec:     72.56
python-uwsgi               Requests/sec:      0.80
python-gunicorn-default    Requests/sec:      0.57
rust-hyper                 Requests/sec:      0.40
nginx                      Requests/sec:       N/A
```
