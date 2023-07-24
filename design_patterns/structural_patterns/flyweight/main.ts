class Address {

    private street: string
    private city: string
    private state: string

    public constructor(street: string, city: string, state: string) {
        this.street = street
        this.city = city
        this.state = state
    }

    public printAddress(): void {
        console.log("Customer address, street:", this.street, "city:", this.city, "state:", this.state)
    }

}

class AddressFactory {

    private static addresses: Address[] = [];

    public static getAddress(address: Address): Address {
        
        var res = AddressFactory.addresses.find((address) => { return address })
        
        if (res == undefined) {
            AddressFactory.addresses.push(address)
            return address
        }

        return res

    }

}

class Customer {

    private name: string
    private address: Address

    public constructor(name: string, address: Address) {
        this.name = name
        this.address = address
    }

    public printCustomer(): void {
        console.log("Customer name:", this.name)
        this.address.printAddress()
    }

}

var padjajaran = AddressFactory.getAddress(new Address("Jl. Padjajaran", "Jakarta", "Indonesia"))
var kencana = AddressFactory.getAddress(new Address("Jl. Kencana", "Bogor", "Indonesia"))
var sandamana = AddressFactory.getAddress(new Address("Jl. Sandamana", "Bogor", "Indonesia"))
var kesasa = AddressFactory.getAddress(new Address("Jl. Kesasa", "Depok", "Indonesia"))

var customer = new Customer("Rival", kencana)
customer.printCustomer()