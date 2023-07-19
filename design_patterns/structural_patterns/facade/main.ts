class SubsystemProduct {
    public newProduct(): void {
        console.log("x-tra shampoo")
    }
}

class SubsystemReport {
    public createReport(): void {
        console.log("report is created!")
    }
}

class SubsystemPromo {
    public createPromo(s: string): void {
        console.log("promo", s, "created!")
    }
    
    public getPromo(): void {
        console.log("promo of x-tra shampoo is only available on july!")
    }
}

class CallCenter {
    private product: SubsystemProduct
    private report: SubsystemReport
    private promo: SubsystemPromo

    constructor() {
        this.product = new SubsystemProduct()
        this.report = new SubsystemReport()
        this.promo = new SubsystemPromo()
    }

    public newProduct(): void {
        this.product.newProduct()
    }

    public createReport(): void {
        this.report.createReport()
    }

    public getPromo(): void {
        this.promo.getPromo()
    }
}

class Manager {
    private promo: SubsystemPromo

    constructor() {
        this.promo = new SubsystemPromo()
    }

    public createPromo(s: string): void {
        this.promo.createPromo(s)
    }
}

var callCenter = new CallCenter()
callCenter.getPromo() // output: promo of x-tra shampoo is only available on july!

var manager = new Manager()
manager.createPromo("buy 1 get 1") // output: promo buy 1 get 1 created!