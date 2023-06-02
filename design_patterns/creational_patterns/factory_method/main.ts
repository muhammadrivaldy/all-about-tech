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

abstract class FactoryGardener {
    abstract createPerson(): Gardener

    public gardener(): void {
        var person = this.createPerson()
        person.maintain()
        person.watering()
    }
}

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
