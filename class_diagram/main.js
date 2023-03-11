"use strict";
exports.__esModule = true;
var example_1 = require("./example");
var mainStore = new example_1.MainStore(new example_1.CoffeeStore("Store Laundry")).storeInterface;
var result = {
    storeName: mainStore.getStoreName(),
    ownerName: mainStore.getOwnerName(),
    employees: mainStore.getEmployees()
};
console.log(JSON.stringify(result));
