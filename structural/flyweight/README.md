# Flyweight Pattern

## Problem
Violation of memory efficiency principles when a large number of similar objects need to be created, leading to excessive memory usage. This commonly occurs in applications that need to manage thousands of fine-grained objects with shared state, where creating separate instances for each object would be prohibitively expensive in terms of memory consumption and performance.

## Solution
Uses the Flyweight Pattern to minimize memory usage by sharing as much data as possible between similar objects. The pattern separates the intrinsic state (shared, immutable data) from the extrinsic state (context-specific, variable data), storing the intrinsic state in flyweight objects that can be shared across multiple contexts. A factory manages flyweight objects, ensuring that they are reused rather than recreated when needed.