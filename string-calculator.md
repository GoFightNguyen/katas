# String Calculator
This is an exercise in coding, refactoring and test-driven development.

Guidelines:
- work incrementally by:
  - using the simplest solutions possible
  - doing one task at a time
  - not reading ahead
- test first
- refactor after each passing test

> Sourced from [Roy Osherove](https://osherove.com/tdd-kata-1) on June 6, 2020 and modified

## Steps
### 1. Create a simple string calculator
Include an `Add` method.

The input is a `string` called `numbers`.
It can be up to two numbers, separated by commas.

The output is an `int`.
It is the sum of the input.

Examples:
- "" => 0
- "1" => 1
- "1,2" => 3

### 2. `Add` handles any amount of numbers