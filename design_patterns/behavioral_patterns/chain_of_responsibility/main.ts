interface Support {
    setNext(h: Support): void
    handle(): void
}

class Bot implements Support {
    private next?: Support = undefined

    public setNext(h: Support): void {
        this.next = h    
    }

    public handle(): void {
        console.log("Answered by the bot for the frequent questions/problems")

        if (this.next != undefined) {
            this.next.handle()
        }
    }
}

class CallCenter implements Support {
    private next?: Support = undefined

    public setNext(h: Support): void {
        this.next = h    
    }

    public handle(): void {
        console.log("Answered by the call center team")

        if (this.next != undefined) {
            this.next.handle()
        }
    }
}

class Engineer implements Support {
    private next?: Support = undefined

    public setNext(h: Support): void {
        this.next = h    
    }

    public handle(): void {
        console.log("Answered by the engineer")

        if (this.next != undefined) {
            this.next.handle()
        }
    }
}

class Customer {
    private support?: Support
    
    public constructor(s?: Support) {
        this.support = s
    }

    public handle(): void {
        console.log("Call the support")
        
        if (this.support != undefined) {
            this.support.handle()
        }

        console.log("Got the solution")
    }
}

var bot = new Bot()
var callCenter = new CallCenter()
var engineer = new Engineer()

bot.setNext(callCenter)
callCenter.setNext(engineer)

var customer = new Customer(bot)
customer.handle()