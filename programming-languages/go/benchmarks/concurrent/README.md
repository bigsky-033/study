# Concurrent

## Book DB

There are four BookDB implementations in this package.

1. With `sync.Map`.
2. With `sync.RWMutex`.
3. With `non buffered channel`, offers sync api.
4. With `buffered channel`, offers async api.

BookDB offers three apis for benchmark.

- Write single instance(Book).
- Read single instance(Book).
- Create snapshot which means it needs to loop safely current objects.

### Test 1 - All operations

#### Test condition

- Read:Write:Snapshot = 20:10:1
- cpu: 4
- execution time: 30 seconds

```
$go test -bench=AllOperations -cpu=4 -benchtime=30s                                                   

goos: linux
goarch: amd64
pkg: github.com/bigsky-033/study/go/benchmarks/concurrent
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkBookDBWithSyncMapAllOperations-4        	   10000	   4174180 ns/op
BenchmarkBookDBWithChannelSyncAllOperations-4    	    4438	   8579380 ns/op
BenchmarkBookDBWithLockAllOperations-4           	    5932	   6191751 ns/op
BenchmarkBookDBWIthChannelAsyncAllOperations-4   	    5588	   6659937 ns/op
PASS
ok  	github.com/bigsky-033/study/go/benchmarks/concurrent	156.021s
```

### Test 2 - Only Read & Write

- Read:Write:Snapshot = 2:1
- cpu: 4
- execution time: 30 seconds

```
$go test -bench=OnlyReadAndWrite -cpu=4 -benchtime=30s

goos: linux
goarch: amd64
pkg: github.com/bigsky-033/study/go/benchmarks/concurrent
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkBookDBWithSyncMapOnlyReadAndWrite-4        	   64257	    552467 ns/op
BenchmarkBookDBWithChannelSyncOnlyReadAndWrite-4    	   13393	   2673528 ns/op
BenchmarkBookDBWithLockOnlyReadAndWrite-4           	   47743	    752244 ns/op
BenchmarkBookDBWIthChannelAsyncOnlyReadAndWrite-4   	   24369	   1506014 ns/op
PASS
ok  	github.com/bigsky-033/study/go/benchmarks/concurrent	199.307s
```
