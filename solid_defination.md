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

A class should have only one reason to change

The statement above means that the class has only one responsibility or single job so every class has a specialty. Like waiters and chefs in the restaurant, they have specific responsibilities to each other. The waiter's responsibility is to gather the request from customers and put it in the note and send it to the kitchen. On the other side, the chef's responsibility is to the food based on the request from the waiter. The waiters and the chefs have different responsibilities, this is an example of _"Single responsibility principle"_.

The benefit of implementing _"Single responsibility"_ is the class will have focus. If you force the class to have more than one responsibility, it will make the developer confused because the class has many abilities. Let's make an example with `Typescript`. We will make a new class with the name _"Employee"_ that handles the employee's detail. This class has many responsibilities and we need to fix it with a _Single responsibility principle_ implementation.

#### Bad Implementation

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

#### Single responsibility implementation

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

Open/Closed principle has a statement `"software entities (classes, modules, functions, etc) should be open for extension, but closed for modification"`, the meaning of that statement is you can extend or copying the behavior of parent class but you cannot modify it.

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
How To Use Open Closed Principle in PHP/Laravel | <https://mohasin-dev.medium.com/how-to-use-open-closed-principal-in-php-laravel-af4fa3b2a1c1>
