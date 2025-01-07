# Benchmarks
Tested with ``AMD Ryzen 9 7940HS 8 Cores 16 Threads`` with 32 GBs RAM and NVME M.2 Gen 4

* badger
```
arch: windows - amd64
engine: badger
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 11.257s    177670 ops/s
get: 4.268s     468592 ops/s

put + get: 15.525s
file size: 242.72MB
peak sys mem: 3.04GB
```

* bbolt
```
arch: windows - amd64
engine: bbolt
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 314.129s   6366 ops/s
get: 1.546s     1293421 ops/s

put + get: 315.675s
file size: 2.00GB
peak sys mem: 451.18MB
```

* goleveldb
```
arch: windows - amd64
engine: goleveldb
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 22.528s    88777 ops/s
get: 7.366s     271510 ops/s

put + get: 29.894s
file size: 714.41MB
peak sys mem: 570.37MB
```

* nutsdb
```
arch: windows - amd64
engine: nutsdb
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 20.128s    99366 ops/s
get: 31.560s    63371 ops/s

put + get: 51.687s
file size: 768.00MB
peak sys mem: 1014.09MB
```

* pebble
```
arch: windows - amd64
engine: pebble
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 9.759s     204937 ops/s
get: 18.187s    109968 ops/s

put + get: 27.946s
file size: 229.03MB
peak sys mem: 538.12MB
```

* pogreb
```
arch: windows - amd64
engine: pogreb
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 54.844s    36466 ops/s
get: 0.325s     6150749 ops/s

put + get: 55.169s
file size: 758.60MB
peak sys mem: 448.68MB
```

* rosedb
```
arch: windows - amd64
engine: rosedb
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 14.647s    136550 ops/s
get: 9.765s     204815 ops/s

put + get: 24.411s
file size: 786.18MB
peak sys mem: 864.77MB
```

* bitcask
```
arch: windows - amd64
engine: bitcask
keys: 2000000
key size: 16-64
value size 128-512
concurrency: 5

put: 38.794s    51554 ops/s
get: 3.040s     657841 ops/s

put + get: 41.834s
file size: 839.22MB
peak sys mem: 3.37GB
```