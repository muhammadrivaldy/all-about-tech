interface Employee {
    selling(): void
}

class EmployeeJohn implements Employee {
    public selling(): void {
        console.log("selling by john")
    }
}

class EmployeeRival implements Employee {
    public selling(): void {
        console.log("selling by rival")
    }
}