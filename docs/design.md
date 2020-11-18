Design
===

The major idea is from `<The ART of Practical Synchronization>` with these changes/optimizations:

1. Using more node types for compressing memory.
2. Using 12 bytes prefix based on CAS 16 bytes(may save more memory).
3. Using pessimistic way to search key avoiding to store key in leaf node.
4. Each node has a pointer to leaf node if has.

## 

