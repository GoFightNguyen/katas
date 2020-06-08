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
- `""` => `0`
- `"1"` => `1`
- `"1,2"` => `3`

### 2. `Add` handles any amount of numbers

### 3. Support new line delimiters too
Examples:
- `"1\n2\n3"` => `6`
- `"1\n2,3"` => `6`

The following is invalid.
You do not need to prove it - just clarifying.
`"1,\n2"`

### 4. Support different delimiters
To change a delimiter, the beginning of the string will contain a separate line like this `//[delimiter]\n`.

Example:
- `"//;\n1;2"` => `3`

All existing scenarios are still supported.
In other words, if a delimiter is not specified then commas and new lines are the defaults.

### 5. No negatives
Calling `Add` with a negative number errors with `negatives not allowed` and the negative that was passed.
If there are multiple negatives then include all of them in the error.

### 6. Ignore numbers bigger than 1000

### 7. Allow delimiters of any length
Delimiters can be of any length with the following format: `//[delimiter]\n`.

Example:
- `"//[***]\n1***2***3"` => `6`

### 8. Allow multiple custom delimiters
Allow multiple delimiters like this: `//[delim1][delim2]\n`.

Example:
- `"//[*][%]\n1*2%3"` => `6`