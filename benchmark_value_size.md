[zhouxinchi@aliyun micro-service-go]$ redis-benchmark -t set,get -n 10000 -d 10
====== SET ======
  10000 requests completed in 0.12 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

94.07% <= 1 milliseconds
99.51% <= 2 milliseconds
100.00% <= 2 milliseconds
81300.81 requests per second

====== GET ======
  10000 requests completed in 0.10 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.80% <= 1 milliseconds
100.00% <= 1 milliseconds
103092.78 requests per second

