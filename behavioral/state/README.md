# State Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) and Open/Closed Principle (OCP) when an object's behavior changes based on its internal state. Using conditional statements (if/else or switch) to manage state transitions leads to complex, hard-to-maintain code that must be modified whenever new states or transitions are added.

## Solution
Uses the State Pattern to allow an object to alter its behavior when its internal state changes. The object will appear to change its class. This pattern encapsulates state-specific behavior into separate state classes and delegates behavior to the current state object, making it easier to add new states and transitions without modifying existing code.