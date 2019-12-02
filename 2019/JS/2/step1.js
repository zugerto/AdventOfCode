"use strict";

var files = require('../util/files');

function solve(program){
    program = program.split(',').map(Number);
    program[1] = 12;
    program[2] = 2;

    return run(program);
}

function run(program){

    const ADD = 1;  
    const MULTIPLY = 2;
    const STOP = 99;

    for (let i = 0; i < program.length; i += 4) {
        let op = program[i];
        let num1 = program[i + 1];
        let num2 = program[i + 2];
        let dest = program[i + 3];

        if (op === ADD) {
            program[dest] = program[num1] + program[num2];
        } else if (op === MULTIPLY) {
            program[dest] = program[num1] * program[num2];
        } else if (op === STOP) {
            break;
        }
    }

    return program[0];
}

console.assert(run([1,0,0,0,99]) == 2);
console.assert(run([2,3,0,3,99]) == 2);
console.assert(run([2,4,4,5,99,0]) == 2);
console.assert(run([1,1,1,4,99,5,6,0,99]) == 30);

//9706670
files.processWithResult('input.txt', solve, console.log);