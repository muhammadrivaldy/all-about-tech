# SOLID Definition

![Status](https://badgen.net/badge/status/completed/green) ![Versions](https://badgen.net/badge/version/v1.0.0/cyan)

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

A class should have only one reason to change.

The statement above means that the class has only one responsibility or single job so every class has a specialty. Like waiters and chefs in the restaurant, they have specific responsibilities to each other. The waiter's responsibility is to gather the request from customers and put it in the note and send it to the kitchen. On the other side, the chef's responsibility is to the food based on the request from the waiter. The waiters and the chefs have different responsibilities, this is an example of _"Single responsibility principle"_.

The benefit of implementing _"Single responsibility"_ is the class will have focus. If you force the class to have more than one responsibility, it will make the developer confused because the class has many abilities. Let's make an example with `Typescript`. We will make a new class with the name _"Employee"_ that handles the employee's detail. This class has many responsibilities and we need to fix it with a _"Single responsibility principle"_ implementation.

#### Before implement single responsibility

``` ts
// This class is too general
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
// This class focus on the employee's detail
class EmployeeDetail {

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

// This class focus on the employee's salary
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

Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.

This principle ensures we do not change the existing code when we want to add a new feature, the alternative we can extends the code and make a new code with the parent code. Let's say you have a recipe for coffee and your employee wants to make a new variant by adding a topping to the coffee. The employee doesn't need to modify your coffee's recipe, he only needs to use your current recipe and add a new topping there. So, your cafe will have a new coffee variant without changing the main recipe.

Sometimes codes are dependent on each other, if you change the code it will impact to the other codes. We can solve it with the _"Open/Closed principle"_. So, the benefit of implementing the _"Open/Closed principle"_ is you don't need to worry your change will impact the existing code. Because you only extend the existing code and make a new code with it.

#### Before implement open/closed

```ts
class Coffee {

    topping: string

    constructor(topping: string) { this.topping = topping }

    public createCoffee(): string {

        var coffee = "base coffee"

        if (this.topping == "chocolate") {
            coffee += " and put chocolate as a coffee's topping"
        } else if (this.topping == "cheese") {
            coffee += " and put cheese as a coffee's topping"
        }

        return coffee

    }

}
```

#### Open/Closed implementation

```ts
class Coffee {
    public createCoffee(): string {
        return "base coffee"
    }
}

class CoffeeChocolate extends Coffee {
    public createCoffee(): string {
        return super.createCoffee() + " and put chocolate as a coffee's topping"
    }
}

class CoffeeCheese extends Coffee {
    public createCoffee(): string {
        return super.createCoffee() + " and put cheese as a coffee's topping"
    }
}
```

### Liskov Substitution Principle

Derived or child classes must be substitutable for their base or parent classes.

This principle was introduces by Barbara Liskov in 1987. The focus on this principle is make the object can be replaceable by other object with the same type. So the high-class doesn't need to know which child-class will be run inside of his class. So, the high-class still can run with the same behavior.

Let's make an example. We have a store with 2 employees and the employees will take turns looking after the store. The store can still running with one of them and the store doesn't need to know who person looks after it.

#### Liskov substitution implementation

```ts
interface Employee {
    selling(): void
}

class EmployeeJohn implements Employee {
    public selling(): void {
        console.log("selling by john")
    }
}

class EmployeeRival implements Employee {
    public selling(): void {
        console.log("selling by rival")
    }
}
```

### Interface Segregation Principle

Do not force any client to implement an interface which is irrelevant to them.

It is similar to the single responsibility, but the difference this approach is applies to the interface instead of classes. This principle focuses on separating the large interface into the small one. We need to avoid large interfaces to make the interface has one focus of responsibility.

Let's say a seller has a pond with many fishes. If the buyer wants to buy a specific fish, it will be hard for the seller & buyer to searching a specific fish in the pond. To make it easier in the future, the seller needs to make other ponds and put every fish in a specific pond. Don't put two different fish in one pond.

#### Before implement interface segregation

```ts
interface Pond {
    catfish(): void
    swordfish(): void
    salmon(): void
}

class Fish implements Pond {

    public catfish(): void {
        // do something
    }

    public swordfish(): void {
        // do something
    }

    public salmon(): void {
        // do something
    }

}
```

#### Interface segregation implementation

```ts
interface PondBigFish {
    catfish(): void
    swordfish(): void
}

interface PondSmallFish {
    salmon(): void
}

class BigFish implements PondBigFish {

    public catfish(): void {
        // do something
    }

    public swordfish(): void {
        // do something
    }

}

class SmallFish implements PondSmallFish {

    public salmon(): void {
        // do something
    }

}

```

### Dependency Inversion Principle

The principle of dependency inversion refers to the decoupling of software modules.

<!-- First of all, _"Dependency Injection"_ between _"Dependency Inversion"_ is different. You can check the below code to know the different of each other: -->

First of all, _"Dependency Injection"_ between _"Dependency Inversion"_ is different. I will show you the difference later and let's focus on the Dependency Inversion's explanation.

This principle focuses on how we can replace the object with another object without making the high-level class getting crash and the high-level class doesn't have a dependency on the low-level class. Both of them will depend on the interface or abstract class. Lemme give you an example the implementation of both of them.

#### Dependency injection example

```ts
class Machine1500CC {
    public startEngine(): void {
        console.log("start engine 1500 cc")
    }
}

class Machine2000CC {
    public startEngine(): void {
        console.log("start engine 2000 cc")
    }
}

class Car {

    constructor(private machine: Machine1500CC) {
        if (machine instanceof Machine1500CC == false) {
            throw new Error("the Car is only able to receive Machine1500CC")
        }
    }

    public startEngine(): void {
        this.machine.startEngine()
    }
    
}

var car = new Car(new Machine1500CC);
car.startEngine() // output "start engine 1500 cc"

var car = new Car(new Machine2000CC);
car.startEngine() // will get an error because the parameter's type of Car is Machine1500CC
```

#### Dependency inversion example

```ts
interface Machine {
    startEngine(): void
}

class Machine1500CC implements Machine {
    public startEngine(): void {
        console.log("start engine 1500 cc")
    }
}

class Machine2000CC implements Machine {
    public startEngine(): void {
        console.log("start engine 2000 cc")
    }
}

class Car {

    constructor(private machine: Machine) { 
        if (!(machine as Machine)) {
            throw new Error("the Car can only receive an object that implements the Machine interface");
        }
    }
    

    public startEngine(): void {
        this.machine.startEngine()
    }
    
}

var car = new Car(new Machine1500CC);
car.startEngine() // output "start engine 1500 cc"

var car = new Car(new Machine2000CC);
car.startEngine() // output "start engine 2000 cc"
```

## Conclusion

Honestly, implementing SOLID principles will take more effort. On the other side, it will make your code easy to maintain & upgrade. All of the principles have a relation to each other, you cannot just implement only one of them.

## Reference

Title | URL
---|---
SOLID Principle in Programming: Understand With Real Life Examples | <https://www.geeksforgeeks.org/solid-principle-in-programming-understand-with-real-life-examples/>
SOLID (Software Design Principles) | <https://www.techtarget.com/whatis/definition/SOLID-software-design-principles>
SOLID Principles for Programming and Software Design | <https://www.freecodecamp.org/news/solid-principles-for-programming-and-software-design/>
A Solid Guide to SOLID Principles | <https://www.baeldung.com/solid-principles>
SOLID Design Principles Explained: Dependency Inversion Principle with Code Examples | <https://stackify.com/dependency-inversion-principle/>
How To Use Open Closed Principle in PHP/Laravel | <https://mohasin-dev.medium.com/how-to-use-open-closed-principal-in-php-laravel-af4fa3b2a1c1>
Applying SOLID principles to TypeScript | <https://blog.logrocket.com/applying-solid-principles-typescript/>
