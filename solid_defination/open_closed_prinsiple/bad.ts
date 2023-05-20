class Coffee {

    topping: string

    constructor(topping: string) { this.topping = topping }

    public createCoffee(): string {

        var coffee = "base coffee"

        if (this.topping == "chocolate") {
            coffee += " and put chocolate as a coffee's topping"
        }

        return coffee

    }

}