abstract class Coupes {
    abstract makeCoupes(): void
}

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

abstract class Crossovers {
    abstract makeCrossovers(): void
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

interface AbstractFactory {
    coupes(): Coupes
    crossovers(): Crossovers
}

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
        return new NissanCoupes    
    }

    public crossovers(): Crossovers {
        return new NissanCrossovers
    }
}

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