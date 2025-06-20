# Decorator Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) and Open/Closed Principle (OCP) when extending object behavior. Subclassing leads to an explosion of classes for each combination of behaviors, and modifying existing classes to incorporate new functionality risks introducing bugs in already tested code.

## Solution
Uses the Decorator Pattern to attach additional responsibilities to objects dynamically. Decorators provide a flexible alternative to subclassing for extending functionality, allowing behaviors to be added or removed at runtime. This pattern creates a set of decorator classes that wrap the original class and provide additional functionality while keeping the interface consistent.