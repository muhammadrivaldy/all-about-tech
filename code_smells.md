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

In general, code smells is not an issue. Even your code is have code smells, your code still working well. If we talk about "the definition of code smells", it will be refer to the fact that your code is not maintainable and not clean enough. Code smell is just an indicator that your code is not following the standards of a good code.

## All kind of code smells

There are several code smells that you should know. Let's define code smells more clearly. I listed code smells below based on <https://refactoring.guru>

### Bloaters

Bloaters are when code, method, and classes get bigger. They make your code hard to understand by others. The bloaters are covers:

- **Long method**: Too many lines of code in there
- **Large class**: Too many fields, method, or lines of code
- **Primitive obsession**: Use primitive attribute rather than complex object for specific variable (phone number, currency, and etc)
- **Long parameter**: More than three parameters for methods
- **Data clumps**: Different parts of the code using identical variables often

### Object-orientation abusers

Object-orientation abusers are when the coder are not following the OOP principles. The object-orientation abusers are covers:

- **Switch statements**: You have a complex `switch` or `if` statements
- **Temporary field**: Define a variable in the class scope but only used by one method
- **Refused bequest**: Subclass uses only some methods and properties, or modify the behavior of its parent
- **Alternative classes with different interfaces**: Two classes are identical but have different method names

### Change preventers

Change preventers are mean that if you want to change a function in one place, you need to put more effort to change the other files too. The change preventers are covers:

- **Divergent change**: Too many changes in a class that maybe not really related to the class
- **Shotgun surgery**: Making a bit change but give a big impact (to other codes)
- **Parallel inheritance hierarchies**: An inheritance must depend on another inheritance

### Dispensables

Dispensables are mean the code have unnecessary code and do not have relation with the goal of class or function. The dispensables are covers:

- **Comments**: A method is filled with explanatory comments
- **Duplicate code**: Two or more code seems identical
- **Lazy class**: Takes more effort to understand & maintain classes
- **Data class**: A class doesn't have method enough to operate itself
- **Dead code**: A variable, parameter, field, method or class is no longer used
- **Speculative generality**: There are unused class, method, field or parameter

### Couplers

Couplers are mean the code have too much coupling between classes. The couplers are covers:

- **Feature envy**: A method accesses another object more than its own data
- **Inappropriate intimacy**: A class uses internal fields and methods of another class
- **Message chains**: When a client requests an object, that object request to another one, and so on
- **Middle man**: A class perform only one action and delegating work to another class

## Reference

Title | URL
---|---
Code Smells | <https://refactoring.guru/refactoring/smells>
Code Smells by Martin Fowler | <https://martinfowler.com/bliki/CodeSmell.html>
Code Smells: What Are They And How Can I Prevent Them? | <https://linearb.io/blog/what-is-a-code-smell#what-is-a-code-smell>
Fixing and Avoiding the Message Chain Code Smell | <https://linearb.io/blog/message-chains-code-smell>
Primitive Obsession — A Code Smell that Hurts People the Most | <https://medium.com/the-sixt-india-blog/primitive-obsession-code-smell-that-hurt-people-the-most-5cbdd70496e9>
Code Smells | <https://code-smells.com>
