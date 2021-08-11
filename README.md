# Benchmark Results on Raspberry Pi node

+-------------+------+------+-----------------+-----------------+-----------------+-----------------+------------------+
| Name        | Succ | Fail | Send Rate (TPS) | Max Latency (s) | Min Latency (s) | Avg Latency (s) | Throughput (TPS) |
|-------------|------|------|-----------------|-----------------|-----------------|-----------------|------------------|
| silomonitor | 1650 | 0    | 200.2           | 0.61            | 0.06            | 0.27            | 198.6            |
+-------------+------+------+-----------------+-----------------+-----------------+-----------------+------------------+

2021.08.11-09:17:48.993 info  [caliper] [round-orchestrator] 	Finished round 1 (silomonitor) in 10.408 seconds
2021.08.11-09:17:48.993 info  [caliper] [monitor.js] 	Stopping all monitors
2021.08.11-09:17:48.994 info  [caliper] [report-builder] 	### All test results ###
2021.08.11-09:17:48.994 info  [caliper] [report-builder] 	
+-------------+------+------+-----------------+-----------------+-----------------+-----------------+------------------+
| Name        | Succ | Fail | Send Rate (TPS) | Max Latency (s) | Min Latency (s) | Avg Latency (s) | Throughput (TPS) |
|-------------|------|------|-----------------|-----------------|-----------------|-----------------|------------------|
| silomonitor | 1650 | 0    | 200.2           | 0.61            | 0.06            | 0.27            | 198.6            |
+-------------+------+------+-----------------+-----------------+-----------------+-----------------+------------------+


test:
    name: silomonitor
    description: test benchmark
    workers:
      type: local
      number: 10
    rounds:
      - label: silomonitor
        description: silomonitor chancode performance evaluation
        txDuration: 10
        rateControl: 
          type: fixed-load
          opts:
            transactionLoad: 200
        workload:
          module: workload/silomonitor.js
          arguments:
            assets: 1
            contractId: silomonitor
