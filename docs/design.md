Design
===

ART(Adaptive Radix Tree) is designed for main memory index:

1. Saving memory space
2. Fast insert & search
3. Concurrency friendly
4. CPU Cache & pipeline friendly
5. Dynamic structure for user-defined

The main idea is from [`<The Adaptive Radix Tree: ARTful Indexing for Main-Memory Databases>`](https://db.in.tum.de/~leis/papers/ART.pdf).

The implementation is based on [UncP/aili](https://github.com/UncP/aili/tree/master/art) with these optimizations:

1. Compress data structure, removing useless fields
2. More APIs/features
3. Add more types of inner nodes for saving memory
