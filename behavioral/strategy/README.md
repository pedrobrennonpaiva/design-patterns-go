# Strategy Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) and Open/Closed Principle (OCP) by having a class that contains multiple algorithms and must be modified whenever a new algorithm is added. This leads to code that is difficult to maintain and extend, particularly when dealing with conditional logic that selects between different behaviors or algorithms.

## Solution
Uses the Strategy Pattern to define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it, allowing you to select the appropriate algorithm at runtime.