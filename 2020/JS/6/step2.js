"use strict";

var files = require('../util/files');

const solve = (rows) => { 
    var numberOfGroupAnswers = 0;
    var groupAnswers = new Set([...'abcdefghijklmnopqrstuvwxyz']);
    
    var lineNumber = 0;
    while (lineNumber < rows.length){
        const answers = rows[lineNumber];

        if (answers === ''){
            numberOfGroupAnswers = numberOfGroupAnswers + groupAnswers.size;
            groupAnswers = new Set([...'abcdefghijklmnopqrstuvwxyz']);
        } else {
            var lineAnswers = new Set(answers);
            groupAnswers = new Set(
                [...groupAnswers].filter(x => lineAnswers.has(x)));
        }

        lineNumber++;
    }

    return numberOfGroupAnswers + groupAnswers.size;
}

console.assert(solve([
    'abc',
    '',
    'a',
    'b',
    'c',
    '',
    'ab',
    'ac',
    '',
    'a',
    'a',
    'a',
    'a',
    '',
    'b',
   ])==6)

files.processLinesWithResult('input.txt', solve, console.log);