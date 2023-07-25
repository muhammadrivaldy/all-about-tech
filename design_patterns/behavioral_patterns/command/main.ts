class User {
    private static users: Map<string, string> = new Map<string, string>()

    public static addUser(name: string, email: string, password: string): void {
        this.users.set(email + password, name)
    }

    public static getUserName(email: string, password: string): string | undefined {
        return this.users.get(email+password)
    }
}

class Login {
    public operation(email: string, password: string): void {
        let name = User.getUserName(email, password)

        if (name != undefined) {
            console.log("Success login, the user name is " + name)
        } else {
            console.log("The user doesn't exists!")
        }
    }
}

class Registration {
    public operation(name: string, email: string, password: string): void {
        User.addUser(name, email, password)
        console.log("Success register the user!")
    }
}

interface Command {
    execute(): void
}

class LoginCommand implements Command {
    private receiver: Login
    private email: string
    private password: string
    
    public constructor(receiver: Login, email: string, password: string) {
        this.receiver = receiver
        this.email = email
        this.password = password
    }

    public execute(): void {
        this.receiver.operation(this.email, this.password)
    }
}

class RegistrationCommand implements Command {
    private receiver: Registration
    private name: string
    private email: string
    private password: string
    
    public constructor(receiver: Registration, name: string, email: string, password: string) {
        this.receiver = receiver
        this.name = name
        this.email = email
        this.password = password
    }

    public execute(): void {
        this.receiver.operation(this.name, this.email, this.password)
    }
}

class Invoker {
    private command?: Command

    public setCommand(command: Command): void {
        this.command = command
    }

    public executeCommand(): void {
        if (this.command != undefined) {
            this.command.execute()
        }
    }
}

var loginOperation = new Login()
var registrationOperation = new Registration()

var loginCommand = new LoginCommand(loginOperation, "rivaldy@gmail.com", "ThisIsPassword")
var registrationCommand = new RegistrationCommand(registrationOperation, "Rival", "rivaldy@gmail.com", "ThisIsPassword")

var invoker = new Invoker()
invoker.setCommand(registrationCommand)
invoker.executeCommand()
invoker.setCommand(loginCommand)
invoker.executeCommand()