# Poc share value between go thread


This repository is only for studying variable sharing between go threads.

The function here is to simulate a setinterval (from js), where it will retrieve the value from an api, and make it available for reading by another thread where the relationship between set and get is 1:n.

In the cases of ref and channel, it is intentional for unit tests to return data race


To run the tests

``` shell
make test
```

To run the benchmark
``` shell
make benchmark
```