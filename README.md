# kv-bench
Fork of https://github.com/lotusdblabs/kvstore-bench which is a fork of https://github.com/akrylysov/pogreb-bench for our own benchmarking of database engines.

## Help
```
  -alsologtostderr
        log to standard error as well as files
  -c int
        number of concurrent goroutines (default 1)
  -compact
        write keys twice and run compaction after
  -e string
        database engine name.
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -log_link string
        If non-empty, add symbolic links in this directory to the log files
  -logbuflevel int
        Buffer log messages logged at this level or lower (-1 means don't buffer; 0 means buffer INFO only; ...). Has limited applicability on non-prod platforms.
  -logtostderr
        log to standard error instead of files
  -maxk int
        maximum key size (default 64)
  -maxv int
        maximum value size (default 512)
  -mink int
        minimum key size (default 16)
  -minv int
        minimum value size (default 128)
  -n int
        number of keys (default 100000)
  -p string
        database path
  -profile string
        enable profile. cpu, mem, block or mutex
  -stderrthreshold value
        logs at or above this threshold go to stderr (default 2)
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```

## Supported Database Engines
* badger
* bbolt
* goleveldb
* lotusdb
* nutsdb
* pebble
* pogreb
* rosedb
* bitcask

## Usage
```bash
./kv-bench -c 5 -e lotusdb -n 2000000 -p /tmp/lotusdb -profile mem
```
or
```bash
go run . -c 5 -e lotusdb -n 2000000 -p /tmp/lotusdb -profile mem
```