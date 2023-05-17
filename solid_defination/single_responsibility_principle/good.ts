// This class focus on the employee's detail
class EmployeeDetail_GoodExample {

    private name: string;
    private address: string;

    constructor(name: string, address: string) {
        this.name = name;
        this.address = address;
    }

    public getEmployeeName(): string {
        return this.name;
    }

    public getEmployeeAddress(): string {
        return this.address;
    }

}

// This class focus on the employee's salary
class EmployeeSalaries_GoodExample {
    
    private attendance: number;
    private salary: number;

    constructor(attendance: number, salary: number) {
        this.attendance = attendance;
        this.salary = salary;
    }

    public getTotalSalaries(): number {
        return this.salary * this.attendance;
    }

}