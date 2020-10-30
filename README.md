# artree
[Adaptive Radix Tree](https://db.in.tum.de/~leis/papers/ART.pdf) in Go.

## Why Adaptive Radix Tree

Main memory capacities have grown up to a point
where most databases fit into RAM. For main-memory database
systems, index structure performance is a critical bottleneck.
Traditional in-memory data structures like balanced binary
search trees are not efficient on modern hardware, because they
do not optimally utilize on-CPU caches. Hash tables, also often
used for main-memory indexes, are fast but only support point
queries.

Adaptive Radix Tree overcomes these shortcomings.


