# ART

Designed for main memory index(Read optimization write exclusion):

1. Read optimization
2. Write exclusion
3. CPU Cache & Pipeline friendly
4. Saving memory space

# Introduction

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

You could find more details in [`<The Adaptive Radix Tree: ARTful Indexing for Main-Memory Databases>`](https://db.in.tum.de/~leis/papers/ART.pdf).

ART(Adaptive Radix Tree) is designed for main memory index:

## Why Adaptive Radix Tree is not popular

It was proposed for the first time in 2010, it's too new for a memory index structure.

There is no mature implementation in present.

# Limitation

## Platform

Only supports X86-64:

16bytes value could be accessed atomically.

## Cannot mix pointer and value in inner node

Go's invalid pointer checking will find the invalid pointer(value). Although we can disable it,
it may cause unpredictable collapse.

# Acknowledge

1. ART index, [UncP/aili](https://github.com/UncP/aili/tree/master/art) 

2. ART based on `<The ART of Practical Synchronization>` [flode/ARTSynchronized](https://github.com/flode/ARTSynchronized)

Thanks for their contribution!
