# Design Patterns

![Status](https://badgen.net/badge/status/in%20progress/orange) ![Versions](https://badgen.net/badge/version/v0.0.1/cyan)

## Overview

Design patterns is a solution to solving a common problem in software design. It's not only to solve one case, design patterns are coverage many cases in software design. You can use it as a blueprint, so it will be helping you to develop your code with a pattern. Design patterns are concepts, you cannot copy them into your program without an adjustment.

## Why we need to know?

Design patterns will help you and your team to develop good code and you can know the code from your teammates easily. This approach will help you to create a code that is maintainable, reusable, & readable. Learning these patterns helps unexperienced developers to learn software design in easy way and faster way.

## Various of Design Patterns

### Creational Patterns

The creational patterns are focused on how to create objects and the client doesn't need to know how the logic has been made. This approach will be decoupling the object from the code and make your code more maintainable and extensible.

These patterns will be divided into:

* #### Factory Method

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

* #### Abstract Factory

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

* #### Builder
  
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

* #### Prototype
  
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

* #### Singleton

### Structural Pattern

* #### Adapter

* #### Bridge

* #### Composite

* #### Decorator

* #### Facade

* #### Flyweight

* #### Proxy

### Behavioral Pattern

* #### Chain of Responsibility

* #### Command

* #### Iterator

* #### Mediator

* #### Memento

* #### Observer

* #### State

* #### Strategy

* #### Template Method

* #### Visitor

## Reference

Title | URL
---|---
Chat GPT | <https://chat.openai.com/chat>
Design Patterns | <https://refactoring.guru/design-patterns>
Design Pattern - Overview | <https://www.tutorialspoint.com/design_pattern/design_pattern_overview.htm>
Design Patterns | <https://sourcemaking.com/design_patterns>
