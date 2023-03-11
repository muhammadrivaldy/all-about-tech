import {Employee, MainStore, CoffeeStore} from "./example";

interface Result {
    storeName: string;
    ownerName: string;
    employees: Employee[];
}

var mainStore = new MainStore(new CoffeeStore("Store Laundry")).storeInterface;

var result: Result = {
    storeName: mainStore.getStoreName(),
    ownerName: mainStore.getOwnerName(),
    employees: mainStore.getEmployees()
}

console.log(JSON.stringify(result));