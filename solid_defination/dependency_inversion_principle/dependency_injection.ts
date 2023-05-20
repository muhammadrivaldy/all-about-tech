class Machine1500CC {
    public startEngine(): void {
        console.log("start engine 1500 cc")
    }
}

class Machine2000CC {
    public startEngine(): void {
        console.log("start engine 2000 cc")
    }
}

class Car {

    constructor(private machine: Machine1500CC) {
        if (machine instanceof Machine1500CC == false) {
            throw new Error("the car is only able to receive Machine1500CC")
        }
    }

    public startEngine(): void {
        this.machine.startEngine()
    }
    
}

var car = new Car(new Machine1500CC);
car.startEngine() // output "start engine 1500 cc"

var car = new Car(new Machine2000CC);
car.startEngine() // will get an error because the parameter's type of Car is Machine1500CC