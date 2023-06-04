class Car {
    private chassis: string
    private engineCylinder: number
    private body: string
    private variant: string

    public setChassis(chassis: string): void {
        this.chassis = chassis
    }

    public setEngineCylinder(engineCylinder: number): void {
        this.engineCylinder = engineCylinder
    }

    public setBody(body: string): void {
        this.body = body
    }

    public setVariant(variant: string): void {
        this.variant = variant
    }
}

interface Builder {
    reset(): void
    buildChassis(): any
    buildEngineCylinder(cylinder: number): any
    buildBody(): any
    buildSportVariant(): any
}

class CarBuilder implements Builder {
    private car: Car
    
    public reset(): void {
        this.car = new Car()
    }

    public buildChassis(): any {
        this.car.setChassis("Standard Chassis")
        return this 
    }

    public buildEngineCylinder(engineCylinder: number) {
        this.car.setEngineCylinder(engineCylinder)
        return this
    }

    public buildBody() {
        this.car.setBody("Standard Body")
        return this
    }

    public buildSportVariant() {
        this.car.setVariant("Sport Variant")
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

    public buildStandardCar() {
        this.builder.reset()
        this.builder.buildChassis().
            buildEngineCylinder(3).
            buildBody()
    }

    public buildSportCar() {
        this.builder.reset()
        this.builder.buildChassis().
            buildEngineCylinder(6).
            buildBody().
            buildSportVariant()
    }
}