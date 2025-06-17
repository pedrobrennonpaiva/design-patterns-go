# Factory Method Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) and Open/Closed Principle (OCP) by having a single function responsible for creating different types of cars. Adding new car models requires modifying existing code.

## Solution
Uses the Factory Method Pattern to delegate object creation to subclasses, allowing the addition of new car types without modifying existing client code.