# Composite Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) and Interface Segregation Principle (ISP) when dealing with tree-like object structures. Client code needs to distinguish between simple objects (leaves) and containers (composites), leading to complex conditional logic. This creates inconsistent interfaces across objects in the hierarchy and complicates operations that should work uniformly on both individual objects and compositions.

## Solution
Uses the Composite Pattern to compose objects into tree structures to represent part-whole hierarchies. The pattern lets clients treat individual objects and compositions of objects uniformly by defining a common interface that both simple and composite objects implement. This simplifies client code, eliminates type checking, and allows for recursive composition where composite objects can contain both leaf objects and other composites.