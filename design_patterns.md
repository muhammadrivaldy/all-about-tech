# Design Patterns

![Status](https://badgen.net/badge/status/in%20progress/orange) ![Versions](https://badgen.net/badge/version/v0.0.1/cyan)

## Article Framework

- [Design Patterns](#design-patterns)
  - [Article Framework](#article-framework)
  - [Overview](#overview)
  - [Why we need to know?](#why-we-need-to-know)
  - [Various of Design Patterns](#various-of-design-patterns)
    - [Creational Patterns](#creational-patterns)
      - [Factory Method](#factory-method)
      - [Abstract Factory](#abstract-factory)
      - [Builder](#builder)
      - [Prototype](#prototype)
      - [Singleton](#singleton)
    - [Structural Patterns](#structural-patterns)
      - [Adapter](#adapter)
      - [Bridge](#bridge)
      - [Composite](#composite)
      - [Decorator](#decorator)
      - [Facade](#facade)
      - [Flyweight](#flyweight)
      - [Proxy](#proxy)
    - [Behavioral Pattern](#behavioral-pattern)
      - [Chain of Responsibility](#chain-of-responsibility)
      - [Command](#command)
      - [Iterator](#iterator)
      - [Mediator](#mediator)
      - [Memento](#memento)
      - [Observer](#observer)
      - [State](#state)
      - [Strategy](#strategy)
      - [Template Method](#template-method)
      - [Visitor](#visitor)
  - [Reference](#reference)

## Overview

Design patterns is a solution to solving a common problem in software design. It's not only to solve one case, design patterns are coverage many cases in software design. You can use it as a blueprint, so it will be helping you to develop your code with a pattern. Design patterns are concepts, you cannot copy them into your program without an adjustment.

## Why we need to know?

Design patterns will help you and your team to develop good code and you can know the code from your teammates easily. This approach will help you to create a code that is maintainable, reusable, & readable. Learning these patterns helps unexperienced developers to learn software design in easy way and faster way.

## Various of Design Patterns

### Creational Patterns

The creational patterns are focused on how to create objects and the client doesn't need to know how the logic has been made. This approach will be decoupling the object from the code and make your code more maintainable and extensible.

This catalog will be divided into:

#### Factory Method

![Factory](https://refactoring.guru/images/patterns/diagrams/factory-method/structure.png)

The factory method will help you to make an object with an interface and the client doesn't need to know how the creation logic works. This approach will help you to extend your feature without any pain.

An example, you expect your garden will maintain today and you do not care who person will maintain your garden. Based on that case, you can use the factory method. Because you only expect the garden will maintain by the gardener without thinking about who it is. And the gardener can be handled by A or B persons.

Let's making the code's example based on above case.

**Step 1:** Create the interface first and that interface will implements by A & B persons.

```ts
interface Gardener {
    maintain(): void
    watering(): void
}

class APerson implements Gardener {
    public maintain(): void {
        console.log("A person maintain garden")
    }

    public watering(): void {
        console.log("A person watering garden")
    }
}

class BPerson implements Gardener {
    public maintain(): void {
        console.log("B person maintain garden")
    }

    public watering(): void {
        console.log("B person watering garden")
    }
}
```

**Step 2:** Create the factory

```ts
abstract class FactoryGardener {
    abstract createPerson(): Gardener

    public gardener(): void {
        var person = this.createPerson()
        person.maintain()
        person.watering()
    }
}
```

**Step 3:** Create the creator class of A & B persons

```ts
class CreatorAPerson extends FactoryGardener {
    public createPerson(): Gardener {
        return new APerson
    }
}

class CreatorBPerson extends FactoryGardener {
    public createPerson(): Gardener {
        return new BPerson
    }
}
```

**Step 4:** Create the client class

```ts
class Client {

    constructor(person: string) {

        var gardener: FactoryGardener;

        if (person == "A") {
            gardener = new CreatorAPerson
        } else if (person == "B") {
            gardener = new CreatorBPerson
        } else {
            throw Error("Undefined")
        }

        gardener.gardener()

    }
}

new Client("A")
```

#### Abstract Factory

![AbstractFactory](https://refactoring.guru/images/patterns/diagrams/abstract-factory/structure-2x.png)

The abstract factory is a common pattern that you implement in the project. This pattern can be called a factory of factories because the focus of this pattern is to produce families of related objects without specifying their concrete classes.

An example, you have the Nissan factory and I have the Toyota factory. Both of us want to make a new car with similar types. The types of cars that we want to create involve Coupes & Crossovers.

Let's making the code's example based on above case.

**Step 1:** Create the abstract product, it means I create abstract class for Coupes & Crossovers. At this moment, it's okay if you want to use interface rather than abstract class.

```ts
abstract class Coupes {
    abstract makeCoupes(): void
}

abstract class Crossovers {
    abstract makeCrossovers(): void
}
```

**Step 2:** Create the concrete product based on the abstract or interface that you create in the step 1. The class will define the method based on that.

```ts
class ToyotaCoupes extends Coupes {
    public makeCoupes(): void {
        console.log("Toyota make coupes")
    }
}

class NissanCoupes extends Coupes {
    public makeCoupes(): void {
        console.log("Nissan make coupes")
    }
}

class ToyotaCrossovers extends Crossovers {
    public makeCrossovers(): void {
        console.log("Toyota make crossovers")
    }
}

class NissanCrossovers extends Crossovers {
    public makeCrossovers(): void {
        console.log("Nissan make crossovers")
    }
}
```

**Step 3:** Create the abstract factory interface. This interface will implements by the Toyota & Nissan factories.

```ts
interface AbstractFactory {
    coupes(): Coupes
    crossovers(): Crossovers
}
```

**Step 4:** Create the concrete factory. This class will implements the interface from `AbstractFactory` and depends to concrete products.

```ts
class ToyotaFactory implements AbstractFactory {
    public coupes(): Coupes {
        return new ToyotaCoupes()
    }

    public crossovers(): Crossovers {
        return new ToyotaCrossovers()
    }
}

class NissanFactory implements AbstractFactory {
    public coupes(): Coupes {
        return new NissanCoupes()  
    }

    public crossovers(): Crossovers {
        return new NissanCrossovers()
    }
}
```

**Step 5:** And the last step is create a client class that have a dependency inversion into the abstract factory interface.

```ts
class Client {
    private factory: AbstractFactory

    constructor(factory: AbstractFactory) { 
        this.factory = factory
    }

    public makeCoupes(): Coupes {
        return this.factory.coupes()
    }

    public makeCrossovers(): Crossovers {
        return this.factory.crossovers()
    }
}

var client = new Client(new NissanFactory)

var coupes = client.makeCoupes()
coupes.makeCoupes()

var crossovers = client.makeCrossovers()
crossovers.makeCrossovers()
```

#### Builder

![Builder](https://refactoring.guru/images/patterns/diagrams/builder/structure.png)

The builder is a pattern to split your complex object into a step-by-step approach. This pattern suggests extracting the construction to separate objects called builders. So the builder class will independent itself.

An example, if you wanna make an object called a Car. You need to make a step-by-step to complete the car's construction. You need to make a chassis, engine cylinder, body, etc. And other times, you need to make the same object car but with car's variants.

Let's making the code's example based on above case.

**Step 1:** Create a `Car` class, this class is a complex object that we want to solved with the builder pattern.

```ts
class Car {
    public chassis: string
    public engineCylinder: number
    public body: string
    public variant: string

    constructor() {
        this.chassis = ""
        this.engineCylinder = 0
        this.body = ""
        this.variant = ""
    }
}
```

**Step 2:** Create a `Builder` class & interface, this is the most important thing in the builder pattern. Because those class & interface will help a `Car` class to make an object step-by-step.

```ts
interface Builder {
    reset(): any
    buildChassis(): any
    buildEngineCylinder(cylinder: number): any
    buildBody(): any
    buildSportVariant(): any
}

class CarBuilder implements Builder {
    private car: Car

    constructor() {
        this.car = new Car()
    }
    
    public reset(): any {
        this.car = new Car()
        return this
    }

    public buildChassis(): any {
        this.car.chassis = "Standard Chassis"
        return this 
    }

    public buildEngineCylinder(engineCylinder: number) {
        this.car.engineCylinder = engineCylinder
        return this
    }

    public buildBody() {
        this.car.body = "Standard Body"
        return this
    }

    public buildSportVariant() {
        this.car.variant = "Sport Variant"
        return this
    }

    public getResult(): Car {
        return this.car
    }
}
```

**Step 3:** Create a `Director` class, this class has a role as a runner of the builder.

```ts
class Director {
    private builder: Builder

    constructor(builder: Builder) {
        this.builder = builder
    }

    public makeStandardCar() {
        this.builder.reset().
            buildChassis().
            buildEngineCylinder(3).
            buildBody()
    }

    public makeSportCar() {
        this.builder.reset().
            buildChassis().
            buildEngineCylinder(6).
            buildBody().
            buildSportVariant()
    }
}
```

**Step 4:** Create a Client class, this client class will handle the dependency process and running the code.

```ts
class Client {
    constructor() {
        var builder = new CarBuilder()
        var director = new Director(builder)
        director.makeSportCar()
        var result = builder.getResult()
        console.log(JSON.stringify(result))
    }
}

new Client
```

#### Prototype

![Prototype](https://refactoring.guru/images/patterns/diagrams/prototype/structure.png)

The prototype is a pattern that helps you to copy the existing object without any dependence on the classes. Let's say you have an object and you want to copy the object, how do you do that? Maybe you will create another object and copy the attribute one by one. That is bad practice to copy the object. Prototype pattern will helps you to copy the object easily.

Let's making the code's example based on above case.

**Step 1:** Create the interface that includes the clone method, the clone method is mandatory for cloning the object. The clone method is the main thing of the prototype pattern.

```ts
interface PrototypeStudent {
    setName(name: string): void
    setGrade(grade: number): void
    getName(): string
    getGrade(): number
    clone(): PrototypeStudent
}
```

**Step 2:** Implement the interface

```ts
class Student implements PrototypeStudent {
    private name: string
    private grade: number

    constructor(name: string, grade: number) {
        this.name = name
        this.grade = grade
    }

    public setName(name: string): void {
        this.name = name
    }

    public setGrade(grade: number): void {
        this.grade = grade
    }

    public getName(): string {
        return this.name
    }

    public getGrade(): number {
        return this.grade
    }

    public clone(): PrototypeStudent {
        return new Student(this.name, this.grade)
    }
}
```

**Step 3:** This is the final step, the client copies the object without any pain to copy it.

```ts
class Client {
    constructor() {
        var studentA = new Student("student A", 9)
        var studentB = studentA.clone()

        studentB.setName("student B")
        studentB.setGrade(4)

        console.log(JSON.stringify(studentA)) // Output: {"name":"student A","grade":9}
        console.log(JSON.stringify(studentB)) // Output: {"name":"student B","grade":4}
    }
}

new Client
```

#### Singleton

![Singleton](https://refactoring.guru/images/patterns/diagrams/singleton/structure-en.png)

The singleton pattern is focused on making sure you only have one instance and sharing the instance for global access. This pattern is applicable to your program which should have a single instance.

Let's making the code's example based on above case.

```ts
class Database {
    private static instance: Database

    private constructor() {}

    public static getInstance(): Database {
        if (!Database.instance) {
            Database.instance = new Database()
        }
        return Database.instance
    }

    public doSomething() {
        console.log("Doing something...")
    }
}

const databaseInstance = Database.getInstance()
databaseInstance.doSomething()
```

### Structural Patterns

The structural patterns are focus on how classes and objects are composed to form larger structures and relationships. This concept will help you to make your code more flexible and efficient.

This catalog will be divided into:

#### Adapter

![Adapter](https://refactoring.guru/images/patterns/diagrams/adapter/structure-object-adapter.png)

The adapter pattern focuses on helping an object to communicate with incompatible interfaces to collaborate. In the real case example, when you travel from Indonesia to Europe you may get a surprise when trying to charge your phone. The power plug and sockets standards are different. So, you need a plug adaptor to charge your phone.

#### Bridge

![Bridge](https://refactoring.guru/images/patterns/diagrams/bridge/structure-en.png)

The bridge pattern focuses on separating the large class into several hierarchies. The bridge's concept is a combination of the single responsibility and dependency inversion in SOLID principles. Basically, if you already learn `Creational Pattern` you will find any bridge implementation in the code's example.

#### Composite

![Composite](https://refactoring.guru/images/patterns/diagrams/composite/structure-en.png)

The composite pattern focuses on composing objects into tree structures and working with if they are individual objects. This pattern is applicable if your app structure is represented as a tree.

In the real-world example, a company has an organization starting from the boss to the employees. When the boss has an idea that needs to be done, he will give a command to the subordinate and the subordinate will continue the command to employees and the employees will work with it.

#### Decorator

![Decorator](https://refactoring.guru/images/patterns/diagrams/decorator/structure.png)

The decorator pattern provides flexible extending functionality. In general, the concept of this pattern is similar to the builder pattern. If the builder is building the object step by step, in the decorator you create a behavior based on several behaviors. Weird right? I think so hahaha.

An example, you already wear a T-shirt but you still getting cold. After that, you wear a sweater but you still getting cold. After that, you wear a coat and now you're warmer.

#### Facade

![Facade](https://refactoring.guru/images/patterns/diagrams/facade/structure.png)

The facade pattern is an interface that encapsulations and simplifies a complex subsystem. So the client doesn't need to know the complexity of the subsystem relations. By doing this, you can decrease the complex relationship between classes.

An example, you want to ask something about the product of the company. That company gas many products there, so you will asking it to the customer service team. The customer service team will help you to answer your question and you don't need to know all of the products that they have.

#### Flyweight

![Flyweight](https://refactoring.guru/images/patterns/diagrams/flyweight/structure.png)

The flyweight pattern focuses to reduce memory usage by sharing data between objects instead of putting all attributes into one object. We separate the attributes based on 2 things, intrinsic (shared) state & extrinsic (unique) state.

An example, you have a company and the company has several divisions. Every division needs to audit by the financial audit to make sure the cash flow is still healthy. If you hire a financial audit team for all divisions and it will be expensive for the company, the solution is to hire one financial audit team and audit the cash flow one by one.

#### Proxy

![Proxy](https://refactoring.guru/images/patterns/diagrams/proxy/structure.png)

The proxy pattern implementation is look like a facade pattern, a proxy will encapsulate another object and you can add something there without any changes in the object that we want to wrap.

An example, if you want to go to an interview at the company and you need to go to the office. Then, when you arrived there you need to get permission first from the security before you can do an interview. The security looks like a proxy.

### Behavioral Pattern

#### Chain of Responsibility

![Chain of responsibility](https://refactoring.guru/images/patterns/diagrams/chain-of-responsibility/structure.png)

The chain of responsibility pattern is a method to pass a request along a chain of handlers until one of them handles it. Oftentimes, this pattern is used in the middleware process for checking something before the request passes to the last handler like authentication or validation and etc.

An example, if you have a problem with an internet provider and you want to report it to them. The step by step of report will be like this:

- Call the call center
- Answered by the bot for frequent questions/problems
- Answered by the call center team
- Answered by the engineer (the last handler)

Based on that, when you did a report, it will be step by step before you can reach the engineer of them. Let's say the call center can handle your problem, so the engineer doesn't need to answer your question.

#### Command

![Command](https://refactoring.guru/images/patterns/diagrams/command/structure.png)

The command pattern has approach to delay or queue your execution by pass requests as a method arguments. Do you understand? I think no, because me too haha. Honestly, it's a little bit hard to understand it in the first time, but after I take a look the code and repeat the explanation often, I got the context and more clear now.

The main thing is, this pattern encapsulates a request as an object and passes it as an object. And this pattern supports you to revert the execution too. Let's take an example for a clearer explanation.

#### Iterator

![Iterator](https://refactoring.guru/images/patterns/diagrams/iterator/structure.png)

The iterator is a pattern that provides a way to access the elements sequentially without exposing their underlying representation. It will iterate over elements without knowing how the elements were implemented.

An example, if you wanna go someplace that you didn't come to that place before and you don't have any idea how to go there. The iterator will help you in several ways:

- Check the place one by one randomly
- Use virtual maps
- Ask the local guide

You can call all as an iterator.

#### Mediator

![Mediator](https://refactoring.guru/images/patterns/diagrams/mediator/structure.png)

The mediator is a pattern that provides reduce dependencies between objects. In this pattern, if an object wants to communicate with other objects, it should through from mediator first instead direct communicating.

An example, we have 2 trains and both of them want to come to a station. Both of them is impossible to communicate with each other, they need to communicate their arrival with the station. So, the station can handle the track and info to them about which track they can use to come to the station.

#### Memento

![Memento](https://refactoring.guru/images/patterns/diagrams/memento/structure1.png)

The memento is a pattern that focuses on saving and restoring states. With memento, you can save the checkpoint and restore it later. This pattern captures the checkpoint and restore the checkpoint without violating encapsulation.

I think, I don't need to explain the real case. It's clear and we can jump out to the code's example.

#### Observer

![Observer](https://refactoring.guru/images/patterns/diagrams/observer/structure.png)

The observer also known as event-subscriber or listener is a pattern that helps you to define a subscription mechanism to notify multiple objects if any events that happen in the publisher.

An example, you want to buy a new phone from the store. But, the release date is unpredictable. Then, the store suggests to you to subscribe their email notification. Therefore, you will get a notification if the release date is fixed.

#### State

![State](https://refactoring.guru/images/patterns/diagrams/state/structure-en.png?id=38c5cc3a610a201e5bc26a441c63d327)

The state pattern is a method that allowed the object to replace its behavior when its internal state changes. Every state that you want to create should be following the interface. So, every state will have the same orientation even though has different logic in there.

An example, you have a tumbler that includes an oil there. After that you want to replace the state (oli) with water, you can put all of the water that you have into tumbler. So, the oil will be gone from the tumbler.

#### Strategy

![Strategy](https://refactoring.guru/images/patterns/diagrams/strategy/structure.png)

The strategy pattern is really similar to the state pattern. The difference is only about the purpose itself. The strategy pattern lets you define a family of algorithms and put it in separate classes and the classes will interchangeable.

The analogy of this pattern is similar to the iterator pattern as well. An example, you can go to the airport with several transportation either using a bicycle, bus, or car. All that transportation will be driving you to the airport.

#### Template Method

#### Visitor

## Reference

Title | URL
---|---
Chat GPT | <https://chat.openai.com/chat>
Design Patterns | <https://refactoring.guru/design-patterns>
Design Pattern - Overview | <https://www.tutorialspoint.com/design_pattern/design_pattern_overview.htm>
Design Patterns | <https://sourcemaking.com/design_patterns>
Design Pattern untuk Pembuatan Perangkat Lunak | <https://www.dicoding.com/blog/design-pattern>
Command Pattern | <https://www.geeksforgeeks.org/command-pattern>
