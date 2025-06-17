# Observer Pattern

## Problem
Violation of the Open/Closed Principle (OCP) when an object needs to notify other objects about its state changes without knowing who they are. This creates tight coupling between the subject and observers, making the system difficult to maintain and extend as adding new subscriber types requires modifying the subject.

## Solution
Uses the Observer Pattern to define a one-to-many dependency between objects so that when one object (the subject) changes state, all its dependents (observers) are notified and updated automatically. This pattern promotes loose coupling between the subject and its observers, allowing observers to be added or removed at runtime without modifying the subject.