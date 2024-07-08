"use strict"

/*
// Objects
const person = {
    name: "Arshdeep",
    surname: "Singh"
};

console.log(person.name, person.surname);

person.name = "Cristiano";
person.surname = "Magrini";
console.log(person.name, person.surname);

person.age = 18;
console.log(person.name, person.surname, person.age);

const prop = prompt("Property?");
console.log(prop, ": ", person[prop]);

for (let i in person) {
    console.log(">", i, ">", person[i]);
}
*/

// Array
const arr = [];
arr.push(1, 2, 3);
arr.unshift(0);
arr.pop();
arr.shift();
arr.slice(1, 2)
arr.splice(1, 2);
console.log(arr);