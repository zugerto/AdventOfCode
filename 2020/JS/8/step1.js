"use strict";

var files = require('../util/files');

const solve = (input) => { 
    var lines = input.trim().split('\n')

    var accumulator = 0;
    var pointer = 0;

    const seen = new Set();
    
    const instructions = lines.map( (line) => {
        const {groups} = line.match(/^(?<instruction>.*) \+?(?<value>-?\d+)$/)
        groups.value = parseInt(groups.value);
        return groups;
    });

    while (!seen.has(instructions[pointer])){
        seen.add(instructions[pointer]);

        const {instruction, value} = instructions[pointer];
        
        switch (instruction) {
            case 'nop':
                pointer++;
                break;
            case 'acc':
                accumulator += value;
                pointer++;
                break;
            case 'jmp':
                pointer += value;
                break;    
            default:
                break;
        }
    }

    return accumulator;
}

console.assert(solve(
    'nop +0\n' +
    'acc +1\n' +
    'jmp +4\n' +
    'acc +3\n' +
    'jmp -3\n' +
    'acc -99\n' +
    'acc +1\n' +
    'jmp -4\n' +
    'acc +6')==5)

files.processWithResult('input.txt', solve, console.log);