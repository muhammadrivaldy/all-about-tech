# SOLID Definition

![Status](https://badgen.net/badge/status/in%20progress/orange) ![Versions](https://badgen.net/badge/version/v0.0.1/cyan)

## Overview

SOLID design principles are promoted by [Robert C. Martin (Uncle Bob)](https://en.wikipedia.org/wiki/Robert_C._Martin). It's best-known design principles in object oriented programming. Those principles will help you to make a good code and the code it will be easier to read, maintenance, upgrade and etc. When you work with an OOP, Solid is a more important thing do you know if you work with a big team.

SOLID is acronym 5 principles of:

- Single Responsibility Principle
- Open/Closed Principle
- Liskov Substitution Principle
- Interface Segregation Principle
- Dependency Inversion Principle

## SOLID Principles

### Single Responsibility Principle

Single responsibility has a statement `“a class should have only one reason to change”`, the meaning of that statement is the class is only have one responsibility or single job. I will give you an example based on restaurant jobs. They have 2 important jobs there "waiter" and "chef". The waiter has a responsibility to take an order from the customer and give the information to the kitchen side, and the chef has a responsibility to make food based on the request from the waiter. They have their own responsibilities and it will make the person more focused on his task rather than one person handling both of job, it will make the person confused.

Often, the developer put all of cases in the one class. That makes the class not readable and not maintainable (too many functions), because the class will have a long code line.

#### Example

We have a class Employee, the class will manage employees information like personal information or salary of employee. We will implement it to the code, and this is the codes:

We have a class Employee, the class will manage employees information like personal information or salary of employee. We have 2 examples (bad & correct) implementations, let's take a look to the codes below:

##### Bad Implementation

It's bad approach because the class `Employee` is manage two things of employees (information & salaries of employee). In single responsibility approach, we should use one class for one thing. Let's continue with the single responsibility implementation.

``` ts
class Employee {

    private name: string;
    private address: string;
    private attendance: number;
    private salary: number;

    constructor(name: string, address: string, attendance: number, salary: number) {
        this.name = name;
        this.address = address;
        this.attendance = attendance;
        this.salary = salary;
    }

    public getEmployeeName(): string {
        return this.name;
    }

    public getEmployeeAddress(): string {
        return this.address;
    }

    public getTotalSalaries(): number {
        return this.salary * this.attendance;
    }

}
```

##### Single Responsibility Implementation

In the previously, one class is manage more than one thing. That is not a principle of `Single responsibility`. Let's take a look to the code below for the correct single responsibility implementation.

``` ts
class Employee {

    private name: string;
    private address: string;

    constructor(name: string, address: string) {
        this.name = name;
        this.address = address;
    }

    public getEmployeeName(): string {
        return this.name;
    }

    public getEmployeeAddress(): string {
        return this.address;
    }

}

class EmployeeSalaries {
    
    private attendance: number;
    private salary: number;

    constructor(attendance: number, salary: number) {
        this.attendance = attendance;
        this.salary = salary;
    }

    public getTotalSalaries(): number {
        return this.salary * this.attendance;
    }

}
```

### Open/Closed Principle

### Liskov Substitution Principle

### Interface Segregation Principle

### Dependency Inversion Principle

## Reference

Title | URL
---|---
SOLID Principle in Programming: Understand With Real Life Examples | <https://www.geeksforgeeks.org/solid-principle-in-programming-understand-with-real-life-examples/>
SOLID (Software Design Principles) | <https://www.techtarget.com/whatis/definition/SOLID-software-design-principles>
SOLID Principles for Programming and Software Design | <https://www.freecodecamp.org/news/solid-principles-for-programming-and-software-design/>
A Solid Guide to SOLID Principles | <https://www.baeldung.com/solid-principles>
SOLID Design Principles Explained: Dependency Inversion Principle with Code Examples | <https://stackify.com/dependency-inversion-principle/>
