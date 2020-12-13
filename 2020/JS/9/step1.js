"use strict";

var files = require('../util/files');

const isValid = (number, numbers) => {
    for (let i = 0; i < numbers.length; i++) {
        for (let j = i + 1; j < numbers.length; j++) {
            if (numbers[i] + numbers[j] === number){
                return true;
            }
        }
    }
    return false;
}

const solveWithPreamble = (preamble, input) => {
    var numbers = input.trim().split('\n').map(x => parseInt(x));

    for (let i = preamble; i < numbers.length; i ++){
        const number = numbers[i];
        if (!isValid(number, numbers.slice(i - preamble, i))){
            return number;
        }
    }
}

const solve = (input) => { 
    return solveWithPreamble(25, input);
}

console.assert(solveWithPreamble(5, 
    '35\n' +
    '20\n' +
    '15\n' +
    '25\n' +
    '47\n' +
    '40\n' +
    '62\n' +
    '55\n' +
    '65\n' +
    '95\n' +
    '102\n' +
    '117\n' +
    '150\n' +
    '182\n' +
    '127\n' +
    '219\n' +
    '299\n' +
    '277\n' +
    '309\n' +
    '576\n')==127)

files.processWithResult('input.txt', solve, console.log);