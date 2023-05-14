class Coffee_BadExample {

    topping: string

    constructor(topping: string) {
        this.topping = topping
    }

    public createCoffee(): void {
        console.log("base coffee")

        if (this.topping == "chocolate") {
            console.log("put chocolate topping to the coffee")
        } else {
            console.log("you're put wrong")
        }
    }

}

var coffeeBad = new Coffee_BadExample("chocolate")
coffeeBad.createCoffee()