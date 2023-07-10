interface PhoneSocket {
    charging(): string
}

class Adaptor implements PhoneSocket {
    private socket: EuropeSocketPlug

    constructor() {
        this.socket = new EuropeSocketPlug
    }

    public charging(): string {
        return this.socket.charging()
    }
}

class EuropeSocketPlug {
    public charging(): string {
        return "sending power"
    }
}

var phone: PhoneSocket
phone = new Adaptor()
console.log(phone.charging()) // output: sending power