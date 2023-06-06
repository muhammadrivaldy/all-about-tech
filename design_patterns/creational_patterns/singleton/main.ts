class Database {
    private static instance: Database

    private constructor() {}

    public static getInstance(): Database {
        if (!Database.instance) {
            Database.instance = new Database()
        }
        return Database.instance
    }

    public doSomething() {
        console.log("Doing something...")
    }
}

const databaseInstance = Database.getInstance()
databaseInstance.doSomething()