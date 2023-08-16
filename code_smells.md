# Code Smells

![Status](https://badgen.net/badge/status/completed/green) ![Versions](https://badgen.net/badge/version/v1.0.0/cyan)

## Table of contents

- [Code Smells](#code-smells)
  - [Table of contents](#table-of-contents)
  - [What is code smells?](#what-is-code-smells)
  - [All kind of code smells](#all-kind-of-code-smells)
    - [Bloaters](#bloaters)
    - [Object-orientation abusers](#object-orientation-abusers)
    - [Change preventers](#change-preventers)
    - [Dispensables](#dispensables)
    - [Couplers](#couplers)
  - [Reference](#reference)

## What is code smells?

In general, code smells is not an issue. Even your code is have code smells, your code still working well. If we talk about "the definition of code smells", it will be refer to the fact that your code is not maintainable and not clean enough.

Code smell is just an indicator that your code is not following the standards of a good code.

## All kind of code smells

There are several code smells that you should know. Let's define code smells more clearly. I listed code smells below based on <https://refactoring.guru>

### Bloaters

Bloaters are when code, method, and classes get bigger. They make your code hard to understand by others. The bloaters are covers:

- Long method
- Large class
- Primitive obsession
- Long parameter
- Data clumps

### Object-orientation abusers

Object-orientation abusers are when the coder are not following the OOP principles. The object-orientation abusers are covers:

- Switch statements
- Temporary field
- Refused bequest
- Alternative classes with different interfaces

### Change preventers

Change preventers are mean that if you want to change a function in one place, you need to put more effort to change the other files too. The change preventers are covers:

- Divergent change
- Shotgun surgery
- Parallel inheritance hierarchies

### Dispensables

Dispensables are mean the code have unnecessary code and do not have relation with the goal of class or function. The dispensables are covers:

- Comments
- Duplicate code
- Lazy class
- Data class
- Dead code
- Speculative generality

### Couplers

Couplers are mean the code have too much coupling between classes. The couplers are covers:

- Feature envy
- Inappropriate intimacy
- Message chains
- Middle man

## Reference

Title | URL
---|---
Code Smells | <https://refactoring.guru/refactoring/smells>
Code Smells by Martin Fowler | <https://martinfowler.com/bliki/CodeSmell.html>
Code Smells: What Are They And How Can I Prevent Them? | <https://linearb.io/blog/what-is-a-code-smell#what-is-a-code-smell>
