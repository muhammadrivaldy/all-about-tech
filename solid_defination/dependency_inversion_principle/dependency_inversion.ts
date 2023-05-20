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
            throw new Error("the car can only receive an object that implements the Machine interface");
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