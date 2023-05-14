class Coffee_GoodExample {

    public createCoffee(): void {
        console.log("base coffee")
    }

}

class CoffeeChocolate_GoodExample extends Coffee_GoodExample {

    public createCoffee(): void {
        super.createCoffee()
        console.log("put chocolate topping to the coffee")
    }
    
}

var coffeeGood = new CoffeeChocolate_GoodExample()
coffeeGood.createCoffee()
