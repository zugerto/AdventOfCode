"use strict";

var files = require('../util/files');

function solve(program){
    program = program.split(',').map(Number);

    let nums = [...Array(100).keys()]

    for (let noun of nums) {
        for (let verb of nums) {
            if (run(program.slice(0),noun,verb) === 19690720) {
                return 100 * noun + verb;
            }
        }
    }
}

function run(program, noun, verb){
    program[1] = noun;
    program[2] = verb;

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

//2552
files.processWithResult('input.txt', solve, console.log);