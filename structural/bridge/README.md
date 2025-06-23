# Bridge Pattern

## Problem
Violation of the Single Responsibility Principle (SRP) and Open/Closed Principle (OCP) when class hierarchies grow exponentially due to combining multiple dimensions of variation. This creates a tight coupling between abstractions and their implementations, making the system difficult to extend and maintain. Each new variation requires creating multiple new classes across the hierarchy.

## Solution
Uses the Bridge Pattern to separate an abstraction from its implementation so that the two can vary independently. The Bridge pattern involves an interface which acts as a bridge between the abstract class and implementor classes, allowing the abstraction hierarchy and the implementation hierarchy to be developed separately without affecting each other. This pattern promotes composition over inheritance, reducing the number of classes and improving flexibility.

Obs: In golang, we don't have abstract class/method (inheritance). To use the pattern correctly, we'll create default struct and add it in parent class.