class CoffeeBase {
    public createCoffee(): string {
        return "base coffee"
    }
}

class CoffeeChocolate extends CoffeeBase {
    public createCoffee(): string {
        return super.createCoffee() + " and put chocolate as a coffee's topping"
    }
}