interface Component {
    execute(): void
}

class BaseDecorator implements Component {

    private wrappee: Component

    constructor(c: Component) {
        this.wrappee = c
    }

    public execute(): void {
        this.wrappee.execute()
    }

}

class Clothes implements Component {

    private clothes: string

    constructor(clothes: string) {
        this.clothes = clothes
    }

    public execute(): void {
        console.log("I wear", this.clothes)        
    }

}

class ExtraClothes extends BaseDecorator {

    private clothes: string

    constructor(c: Component, clothes: string) {
        super(c)
        this.clothes = clothes
    }

    public execute(): void {
        super.execute()
        this.extra()
    }

    public extra(): void {
        console.log("I wear extra:", this.clothes)
    }

}

var tShirt = new Clothes("T-Shirt")
var sweater = new ExtraClothes(tShirt, "Sweater")
var coat = new ExtraClothes(sweater, "Coat")

coat.execute()

/**
 * OUTPUT:
 * 
 * I wear T-Shirt
 * I wear extra: Sweater
 * I wear extra: Coat
 */