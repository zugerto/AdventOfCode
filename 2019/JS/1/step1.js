"use strict";

var files = require('../util/files');

function solve(masses){
    return masses
    .map(mass => {
        return Math.floor(mass / 3) - 2;
    })
    .reduce((sum, required_fuel) => sum + required_fuel, 0)
}

console.assert(solve([12])==2)
console.assert(solve([14])==2)
console.assert(solve([1969])==654)
console.assert(solve([100756])==33583)

files.processLinesWithResult('input.txt', solve, console.log);