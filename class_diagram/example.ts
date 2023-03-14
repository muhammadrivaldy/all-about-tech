export class User {

    private id: number;
    private name: string;

    constructor(id: number, name: string) {
        this.id = id;
        this.name = name;
    }

    public getId(): number {return this.id;}
    public getName(): string {return this.name;}

}

export class Employee extends User {

    private position: string;

    constructor(id: number, name: string, position: string) {
        super(id, name);
        this.position = position;
    }

    public getPosition(): string {return this.position;}

}

export class Owner extends User {

    private balance: number;

    constructor(id: number, name: string, balance: number) {
        super(id, name);
        this.balance = balance;
    }
    
    public getBalance(): number {return this.balance;}

}

interface StoreInterface {
    getStoreName(): string;
    getOwnerName(): string;
    getEmployees(): Employee[];
}

export class CoffeeStore implements StoreInterface {
    
    private storeName: string;
    private owner: Owner;
    private employees: Employee[] = [];

    constructor(storeName: string) {
        this.storeName = storeName;
        this.owner = new Owner(1, "Rival", 1250000);
        this.employees.push(new Employee(2, "Santo", "Manager"));
    }

    public getStoreName(): string {
        return this.storeName;
    }

    public getOwnerName(): string {
        return this.owner.getName();
    }

    public getEmployees(): Employee[] {
        return this.employees;
    }

}

export class CakeStore implements StoreInterface {
    
    private storeName: string;
    private owner: Owner;
    private employees: Employee[] = [];

    constructor(storeName: string) {
        this.storeName = storeName;
        this.owner = new Owner(3, "Steven", 1250000);
        this.employees.push(new Employee(4, "Sudarsono", "Manager"));
    }

    public getStoreName(): string {
        return this.storeName;
    }

    public getOwnerName(): string {
        return this.owner.getName();
    }

    public getEmployees(): Employee[] {
        return this.employees;
    }

}

export class MainStore {

    public storeInterface: StoreInterface;

    constructor(storeInterface: StoreInterface) {
        this.storeInterface = storeInterface;
    }

}