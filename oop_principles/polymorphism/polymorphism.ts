class User {
    public get(): string {
        return "User object"
    }
}

class UserAdmin extends User {
    public get(): string {
        return "User admin object"
    }
}

class UserRegular extends User {
    public get(): string {
        return "User regular object"
    }
}

var user = new User()
console.log(user.get())

var userAdmin = new UserAdmin()
console.log(userAdmin.get())

var userRegular = new UserRegular()
console.log(userRegular.get())