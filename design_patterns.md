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

  An example, you expect your garden will maintain today and you do not care who person will maintain your garden. Based on that case, you can use the factory method. Because you only expect the garden will maintain by the gardener without thinking about who it is. And the gardener can be handled by A or B person.

  Let's making the code's example based on above case.

  **Step 1:** Create the interface first and that interface will implements by A person & B person.

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

  **Step 2:** Create the abstract factory

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
  class Main {

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

  new Main("A")
  ```

* #### Abstract Factory

* #### Builder

* #### Prototype

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
Design Pattern untuk Pembuatan Perangkat Lunak | <https://www.dicoding.com/blog/design-pattern/>
Factory patterns in Go (Golang) | <https://www.sohamkamani.com/golang/2018-06-20-golang-factory-patterns/>
Desing Patterns in Golang: Factory Method | <https://blog.ralch.com/articles/design-patterns/golang-factory-method/>
Design Pattern - Overview | <https://www.tutorialspoint.com/design_pattern/design_pattern_overview.htm>
Design Patterns | <https://sourcemaking.com/design_patterns>
