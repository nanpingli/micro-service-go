
# ======================================
[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> 
127.0.0.1:6379> exit
[zhouxinchi@aliyun workspace]$ 
[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> flushdb
OK
127.0.0.1:6379> INFO memory
# Memory
used_memory:839464
used_memory_human:819.79K
used_memory_rss:4177920
used_memory_rss_human:3.98M
used_memory_peak:14317776
used_memory_peak_human:13.65M
used_memory_peak_perc:5.86%
used_memory_overhead:832134
used_memory_startup:782504
used_memory_dataset:7330
used_memory_dataset_perc:12.87%
total_system_memory:1981448192
total_system_memory_human:1.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:4.98
mem_allocator:jemalloc-3.6.0
active_defrag_running:0
lazyfree_pending_objects:0
127.0.0.1:6379> 
# ======== -n 100000 -d 10 ===============
[zhouxinchi@aliyun workspace]$ redis-benchmark -t get, set -n 100000 -d 10 -q
====== set -n 100000 -d 10 -q ======
  100000 requests completed in 0.98 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1

100.00% <= 1 milliseconds
101729.40 requests per second

[zhouxinchi@aliyun workspace]$ redis-benchmark -t get,set -n 100000 -d 10 -q
SET: 102669.41 requests per second
GET: 103305.79 requests per second

[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> INFO memory
# Memory
used_memory:839592
used_memory_human:819.91K
used_memory_rss:4259840
used_memory_rss_human:4.06M
used_memory_peak:14317776
used_memory_peak_human:13.65M
used_memory_peak_perc:5.86%
used_memory_overhead:832206
used_memory_startup:782504
used_memory_dataset:7386
used_memory_dataset_perc:12.94%
total_system_memory:1981448192
total_system_memory_human:1.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:5.07
mem_allocator:emalloc-3.6.0
active_defrag_running:0
lazyfree_pending_objects:0
127.0.0.1:6379> exit
# ============= -n 100000 -d 1000 ==================
[zhouxinchi@aliyun workspace]$ redis-benchmark -t get,set -n 100000 -d 1000 -q
SET: 102354.15 requests per second
GET: 102986.61 requests per second


[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> INFO memory
# Memory
used_memory:840600
used_memory_human:820.90K
used_memory_rss:4255744
used_memory_rss_human:4.06M
used_memory_peak:14317776
used_memory_peak_human:13.65M
used_memory_peak_perc:5.87%
used_memory_overhead:832206
used_memory_startup:782504
used_memory_dataset:8394
used_memory_dataset_perc:14.45%
total_system_memory:1981448192
total_system_memory_human:1.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:5.06
mem_allocator:jemalloc-3.6.0
active_defrag_running:0
lazyfree_pending_objects:0
127.0.0.1:6379> exit
# ========== -n 100000 -d 10000 ==================
[zhouxinchi@aliyun workspace]$ redis-benchmark -t get,set -n 100000 -d 10000 -q
SET: 89605.73 requests per second
GET: 89847.26 requests per second

[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> INFO memory
# Memory
used_memory:851864
used_memory_human:831.90K
used_memory_rss:4431872
used_memory_rss_human:4.23M
used_memory_peak:14317776
used_memory_peak_human:13.65M
used_memory_peak_perc:5.95%
used_memory_overhead:832206
used_memory_startup:782504
used_memory_dataset:19658
used_memory_dataset_perc:28.34%
total_system_memory:1981448192
total_system_memory_human:1.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:5.20
mem_allocator:jemalloc-3.6.0
active_defrag_running:0
lazyfree_pending_objects:0
127.0.0.1:6379> exit
# ============ -n 100000 -d 100000 ===============
[zhouxinchi@aliyun workspace]$ redis-benchmark -t get,set -n 100000 -d 100000 -q 
SET: 26695.14 requests per second
GET: 12812.30 requests per second


[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> INFO memory
# Memory
used_memory:941976
used_memory_human:919.90K
used_memory_rss:4423680
used_memory_rss_human:4.22M
used_memory_peak:14317776
used_memory_peak_human:13.65M
used_memory_peak_perc:6.58%
used_memory_overhead:832206
used_memory_startup:782504
used_memory_dataset:109770
used_memory_dataset_perc:68.83%
total_system_memory:1981448192
total_system_memory_human:1.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:4.70
mem_allocator:jemalloc-3.6.0
active_defrag_running:0
lazyfree_pending_objects:0
127.0.0.1:6379> exit
[zhouxinchi@aliyun workspace]$ 
# =========== -n 100000 -d 1000000 ============
[zhouxinchi@aliyun workspace]$ redis-benchmark -t get,set -n 100000 -d 1000000 -q
SET: 2028.48 requests per second
GET: 1272.54 requests per second

[zhouxinchi@aliyun workspace]$ redis-cli
127.0.0.1:6379> INFO memory
# Memory
used_memory:1843096
used_memory_human:1.76M
used_memory_rss:5320704
used_memory_rss_human:5.07M
used_memory_peak:81950904
used_memory_peak_human:78.15M
used_memory_peak_perc:2.25%
used_memory_overhead:832206
used_memory_startup:782504
used_memory_dataset:1010890
used_memory_dataset_perc:95.31%
total_system_memory:1981448192
total_system_memory_human:1.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:2.89
mem_allocator:jemalloc-3.6.0
active_defrag_running:0
lazyfree_pending_objects:0
127.0.0.1:6379> 
j
