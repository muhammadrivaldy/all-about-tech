interface PrototypeStudent {
    setName(name: string): void
    setGrade(grade: number): void
    getName(): string
    getGrade(): number
    clone(): PrototypeStudent
}

class Student implements PrototypeStudent {
    private name: string
    private grade: number

    constructor(name: string, grade: number) {
        this.name = name
        this.grade = grade
    }

    public setName(name: string): void {
        this.name = name
    }

    public setGrade(grade: number): void {
        this.grade = grade
    }

    public getName(): string {
        return this.name
    }

    public getGrade(): number {
        return this.grade
    }

    public clone(): PrototypeStudent {
        return new Student(this.name, this.grade)
    }
}

class Client {
    constructor() {
        var studentA = new Student("student A", 9)
        var studentB = studentA.clone()

        studentB.setName("student B")
        studentB.setGrade(4)

        console.log(JSON.stringify(studentA)) // Output: {"name":"student A","grade":9}
        console.log(JSON.stringify(studentB)) // Output: {"name":"student B","grade":4}
    }
}

new Client