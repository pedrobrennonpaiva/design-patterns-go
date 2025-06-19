# Template Method Pattern

## Problem
Violation of the Don't Repeat Yourself (DRY) principle when implementing similar algorithms with shared steps across different classes. This leads to code duplication and maintenance issues, as changes to the algorithm structure require modifications in multiple places. Additionally, it can violate the Open/Closed Principle when algorithm steps need to be extended or customized.

## Solution
Uses the Template Method Pattern to define the skeleton of an algorithm in a base "template" method, deferring some steps to subclasses. This allows subclasses to redefine certain steps of an algorithm without changing the algorithm's structure. In Go, without traditional inheritance, this is implemented using interfaces and composition, where a common structure coordinates the algorithm flow while specific implementations provide the variable steps.

Obs: In golang, we don't have abstract class/method (inheritance). To use the pattern correctly, we'll create default struct and add it in parent class.