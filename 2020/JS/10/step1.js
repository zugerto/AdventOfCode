"use strict";

var files = require('../util/files');

const solve = (input) => { 
    var voltages = input.trim().split('\n').map(x => parseInt(x));

    voltages = [0, ...voltages].sort((a, b) => a - b);
    voltages.push(voltages[voltages.length - 1] + 3);
    
    let voltagesOfOne = 0;
    let voltagesOfThree = 0;

    for (let i = 1; i < voltages.length; i++) {
        const voltageDifference = voltages[i] - voltages[i - 1];
        if (voltageDifference === 1){
            voltagesOfOne++;
        } else if (voltageDifference == 3){
            voltagesOfThree++;
        }
    }

    return voltagesOfOne * voltagesOfThree;
}

console.assert(solve( 
    '16\n' +
    '10\n' +
    '15\n' +
    '5\n' +
    '1\n' +
    '11\n' +
    '7\n' +
    '19\n' +
    '6\n' +
    '12\n' +
    '4\n')==35)

files.processWithResult('input.txt', solve, console.log);