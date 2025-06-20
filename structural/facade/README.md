# Facade Pattern

## Problem
Dealing with complex subsystems often leads to violations of several design principles:
- **Tight coupling**: Client code becomes dependent on the implementation details of subsystems
- **Low cohesion**: Client code mixes business logic with subsystem integration details
- **Interface Segregation Principle violation**: Clients are forced to depend on interfaces they don't use
- **Dependency Inversion Principle violation**: High-level modules directly depend on low-level modules
- **Complexity management**: Client code becomes difficult to understand and maintain due to subsystem complexity

## Solution
Uses the Facade Pattern to provide a simplified, unified interface to a set of interfaces in a subsystem. The Facade defines a higher-level interface that makes the subsystem easier to use by:
- **Reducing coupling**: Client code depends only on the facade, not on complex subsystem classes
- **Improving cohesion**: Subsystem integration details are extracted from client code
- **Following Interface Segregation**: Clients work with a targeted interface specific to their needs
- **Enabling Dependency Inversion**: High-level modules depend on the abstraction provided by the facade
- **Managing complexity**: Hiding subsystem complexity behind a simple interface

The Facade doesn't encapsulate the subsystem classes completely; clients can still use subsystem classes directly when needed, providing flexibility while promoting simplicity for common use cases.