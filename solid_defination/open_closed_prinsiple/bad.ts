class Coffee_BadExample {

    topping: string

    constructor(topping: string) { this.topping = topping }

    public createCoffee(): string {

        if (this.topping == "original") {
            return "base coffee"
        } else if (this.topping == "chocolate") {
            return "base coffee and put chocolate as a coffee's topping"
        }

        throw new Error("we don't have any coffee that you want!")

    }

}