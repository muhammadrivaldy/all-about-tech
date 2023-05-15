class CoffeeBase_GoodExample {
    public createCoffee(): string {
        return "base coffee"
    }
}

class CoffeeChocolate_GoodExample extends CoffeeBase_GoodExample {
    public createCoffee(): string {
        return super.createCoffee() + " and put chocolate as a coffee's topping"
    }
}