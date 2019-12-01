"use strict";

var files = require('../util/files');

function solve(masses){

    let fuelForMass = mass => {
        let fuel = Math.floor(mass / 3) - 2;
        if (fuel <= 0) return 0;
        return fuel + fuelForMass(fuel);
      }

    return masses.reduce((acc, mass) => acc + fuelForMass(mass), 0);
}

console.assert(solve([14])==2)
console.assert(solve([1969])==966)
console.assert(solve([100756])==50346)

files.processLinesWithResult('input.txt', solve, console.log);

