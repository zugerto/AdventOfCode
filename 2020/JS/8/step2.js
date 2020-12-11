"use strict";

var files = require('../util/files');

const solve = (input) => { 
    var lines = input.trim().split('\n')

    const instructions = lines.map( (line) => {
        const {groups} = line.match(/^(?<instruction>.*) \+?(?<value>-?\d+)$/)
        groups.value = parseInt(groups.value);
        return groups;
    });

    for (let i = 0; i < instructions.length; i++) {
        if (instructions[i].instruction === 'nop' || instructions[i].instruction == 'jmp'){
            const copy = JSON.parse(JSON.stringify(instructions));
            copy[i].instruction = instructions[i].instruction == 'nop' ? 'jmp' : 'nop'
         
            var accumulator = 0;
            var pointer = 0;   
            const seen = new Set();
            while (!seen.has(copy[pointer])){
                seen.add(copy[pointer]);

                if (pointer == copy.length){
                    return accumulator;
                }

                const {instruction, value} = copy[pointer];
                
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
        }
    }
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
    'acc +6')==8)

files.processWithResult('input.txt', solve, console.log);