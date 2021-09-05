[zhouxinchi@aliyun micro-service-go]$ redis-benchmark -t set,get -n 100000 -d 10
====== SET ======
  100000 requests completed in 0.96 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

99.95% <= 1 milliseconds
99.95% <= 2 milliseconds
99.98% <= 3 milliseconds
100.00% <= 3 milliseconds
103950.10 requests per second

====== GET ======
  100000 requests completed in 0.96 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
104166.67 requests per second

