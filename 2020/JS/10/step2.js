"use strict";

var files = require('../util/files');

// Dynamic programming example. Recursion with memoization of already computed values.
const countTheNumberOfWaysToComplete = (memorize, voltages, i) => {
    if (i == voltages.length - 1){
        return 1;
    }
    if (memorize.has(i)){
        return memorize.get(i);
    }

    var answer = 0;
    for (let j = i + 1; j < voltages.length; j++) {
        if (voltages[j] - voltages[i] <= 3){
            answer = answer + countTheNumberOfWaysToComplete(memorize, voltages, j);
        }
    }
    memorize.set(i, answer);

    return answer;
}

const solve = (input) => { 
    var voltages = input.trim().split('\n').map(x => parseInt(x));

    voltages = [0, ...voltages].sort((a, b) => a - b);
    voltages.push(voltages[voltages.length - 1] + 3);
    
    var memorize = new Map();
    return countTheNumberOfWaysToComplete(memorize, voltages, 0);
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
    '4\n')==8)

//files.processWithResult('input-sample.txt', solve, console.log);
files.processWithResult('input.txt', solve, console.log);