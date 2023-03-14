"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
exports.__esModule = true;
exports.MainStore = exports.CakeStore = exports.CoffeeStore = exports.Owner = exports.Employee = exports.User = void 0;
var User = /** @class */ (function () {
    function User(id, name) {
        this.id = id;
        this.name = name;
    }
    User.prototype.getId = function () { return this.id; };
    User.prototype.getName = function () { return this.name; };
    return User;
}());
exports.User = User;
var Employee = /** @class */ (function (_super) {
    __extends(Employee, _super);
    function Employee(id, name, position) {
        var _this = _super.call(this, id, name) || this;
        _this.position = position;
        return _this;
    }
    Employee.prototype.getPosition = function () { return this.position; };
    return Employee;
}(User));
exports.Employee = Employee;
var Owner = /** @class */ (function (_super) {
    __extends(Owner, _super);
    function Owner(id, name, balance) {
        var _this = _super.call(this, id, name) || this;
        _this.balance = balance;
        return _this;
    }
    Owner.prototype.getBalance = function () { return this.balance; };
    return Owner;
}(User));
exports.Owner = Owner;
var CoffeeStore = /** @class */ (function () {
    function CoffeeStore(storeName) {
        this.employees = [];
        this.storeName = storeName;
        this.owner = new Owner(1, "Rival", 1250000);
        this.employees.push(new Employee(2, "Santo", "Manager"));
    }
    CoffeeStore.prototype.getStoreName = function () {
        return this.storeName;
    };
    CoffeeStore.prototype.getOwnerName = function () {
        return this.owner.getName();
    };
    CoffeeStore.prototype.getEmployees = function () {
        return this.employees;
    };
    return CoffeeStore;
}());
exports.CoffeeStore = CoffeeStore;
var CakeStore = /** @class */ (function () {
    function CakeStore(storeName) {
        this.employees = [];
        this.storeName = storeName;
        this.owner = new Owner(3, "Steven", 1250000);
        this.employees.push(new Employee(4, "Sudarsono", "Manager"));
    }
    CakeStore.prototype.getStoreName = function () {
        return this.storeName;
    };
    CakeStore.prototype.getOwnerName = function () {
        return this.owner.getName();
    };
    CakeStore.prototype.getEmployees = function () {
        return this.employees;
    };
    return CakeStore;
}());
exports.CakeStore = CakeStore;
var MainStore = /** @class */ (function () {
    function MainStore(storeInterface) {
        this.storeInterface = storeInterface;
    }
    return MainStore;
}());
exports.MainStore = MainStore;
