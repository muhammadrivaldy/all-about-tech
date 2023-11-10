class User {
    public name: string
    public address: string

    constructor(name: string, address: string) {
        this.name = name
        this.address = address
    }
}

class Employee extends User {
    public position: string

    constructor(position: string, name: string, address: string) {
        super(name, address)
        this.position = position
    }

    public printEmployee() {
        console.log(this.name + " is a " + this.position + " and his address is " + this.address)
    }
}

var employee = new Employee("Software Engineer", "Rival", "Jl. Address")
employee.printEmployee()