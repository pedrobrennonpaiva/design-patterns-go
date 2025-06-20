# Adapter Pattern

## Problem
Violation of the Interface Segregation Principle (ISP) and Dependency Inversion Principle (DIP) when integrating components with incompatible interfaces. This occurs when existing classes cannot be modified to conform to interfaces expected by client code, such as when using third-party libraries or legacy code. Without adaptation, this leads to scattered conditional code, tight coupling, and reduced reusability.

## Solution
Uses the Adapter Pattern to convert the interface of a class into another interface that clients expect. The Adapter acts as a wrapper between two incompatible interfaces, allowing them to work together without modifying their source code. This promotes code reuse, separates interface concerns from implementation details, and enables components developed independently to work together seamlessly.