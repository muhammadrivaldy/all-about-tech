import {Employee, MainStore, CoffeeStore, CakeStore} from "./example";

interface Result {
    storeName: string;
    ownerName: string;
    employees: Employee[];
}

var mainCoffeeStore = new MainStore(new CoffeeStore("Store Coffee Abadie")).storeInterface;
var mainCakeStore = new MainStore(new CakeStore("Store Cake Selamanya")).storeInterface;

var resultCoffee: Result = {
    storeName: mainCoffeeStore.getStoreName(),
    ownerName: mainCoffeeStore.getOwnerName(),
    employees: mainCoffeeStore.getEmployees()
}

var resultCake: Result = {
    storeName: mainCakeStore.getStoreName(),
    ownerName: mainCakeStore.getOwnerName(),
    employees: mainCakeStore.getEmployees()
}

console.log(JSON.stringify(resultCoffee));
console.log(JSON.stringify(resultCake));