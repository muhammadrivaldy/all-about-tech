abstract class Info {
    abstract printInformation(): void
}

class UserInfo extends Info {
    public printInformation(): void {
        console.log("This is user information")
    }
}