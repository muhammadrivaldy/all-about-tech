interface Component {
    execute(): void;
}

class Composite implements Component {

    private children: Component[] = []

    public add(c: Component) {
        this.children.push(c)
    }

    public execute() {
        this.children.forEach((child) => child.execute())
    }

}

class EmployeeA implements Component {

    private job: string

    constructor(job: string) {
        this.job = job
    }

    public execute(): void {
        console.log("Employee A:", this.job)
    }

}

class EmployeeB implements Component {

    private job: string

    constructor(job: string) {
        this.job = job
    }

    public execute(): void {
        console.log("Employee B:", this.job)
    }

}

class ManagerA implements Component {

    private job: string

    constructor(job: string) {
        this.job = job
    }

    public execute(): void {
        console.log("Manager A:", this.job)
    }

}

var operation = new Composite()
operation.add(new EmployeeA("Do a research about field"))
operation.add(new EmployeeB("Do a research about stadium"))

var managerial = new Composite()
managerial.add(new ManagerA("Distribute work to subordinates"))
managerial.add(operation)
managerial.execute()

/**
 * OUTPUT:
 * 
 * Manager A: Distribute work to subordinates
 * Employee A: Do a research about field
 * Employee B: Do a research about stadium
 */