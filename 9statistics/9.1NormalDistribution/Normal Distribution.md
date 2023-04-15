# 9.1 Normal Distribution

## Introduction

Toss a coin `ten times`, record  `the number of times it lands on heads` as `one set of tests`.

If you perform `multiple sets of tests`, the data will approach a `normal distribution`.

```bash
$ cd abstractionExplainedz/9statistics/9.1NormalDistribution/toss

$ make test
# go test -v -run='^\QTest_Check_' ./
# === RUN   Test_Check_Toss_And_Count
# === RUN   Test_Check_Toss_And_Count/TossTenTimes
# === RUN   Test_Check_Toss_And_Count/OneThousandCount
# --- PASS: Test_Check_Toss_And_Count (0.00s)
#     --- PASS: Test_Check_Toss_And_Count/TossTenTimes (0.00s)
#     --- PASS: Test_Check_Toss_And_Count/OneThousandCount (0.00s)
# PASS
# ok      github.com/panhongrainbow/abstractionExplainedz/9statistics/9.1NormalDistribution/toss
```

It seems that the simulation results conform to the normal distribution.
