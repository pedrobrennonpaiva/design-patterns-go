# Proxy Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) when objects need additional functionality such as access control, lazy initialization, logging, or caching. Adding these responsibilities directly to the class would make it more complex and harder to maintain. Additionally, the client code may need to work with expensive resources or remote services without being aware of their implementation details.

## Solution
Uses the Proxy Pattern to provide a surrogate or placeholder for another object to control access to it. A proxy maintains a reference to the original object (subject) and implements the same interface, allowing it to be used as a substitute. The proxy can intercept calls to the real object, adding behaviors like access control, caching results, lazy loading, logging, or remote communication, while keeping these concerns separate from the core functionality.