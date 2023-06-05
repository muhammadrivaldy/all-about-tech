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