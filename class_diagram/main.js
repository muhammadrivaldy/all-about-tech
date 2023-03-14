"use strict";
exports.__esModule = true;
var example_1 = require("./example");
var mainCoffeeStore = new example_1.MainStore(new example_1.CoffeeStore("Store Coffee Abadie")).storeInterface;
var mainCakeStore = new example_1.MainStore(new example_1.CakeStore("Store Cake Selamanya")).storeInterface;
var resultCoffee = {
    storeName: mainCoffeeStore.getStoreName(),
    ownerName: mainCoffeeStore.getOwnerName(),
    employees: mainCoffeeStore.getEmployees()
};
var resultCake = {
    storeName: mainCakeStore.getStoreName(),
    ownerName: mainCakeStore.getOwnerName(),
    employees: mainCakeStore.getEmployees()
};
console.log(JSON.stringify(resultCoffee));
console.log(JSON.stringify(resultCake));
