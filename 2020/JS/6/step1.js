"use strict";

var files = require('../util/files');

const solve = (rows) => { 
    var numberOfGroupAnswers = 0;
    var groupAnswers = new Set();
    
    var lineNumber = 0;
    while (lineNumber < rows.length){
        const answers = rows[lineNumber];

        if (answers === ''){
            numberOfGroupAnswers = numberOfGroupAnswers + groupAnswers.size;
            groupAnswers = new Set();
        } else {
            var lineAnswers = new Set(answers);
            groupAnswers = new Set([...groupAnswers, ...lineAnswers]);
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
    ''
   ])==11)

files.processLinesWithResult('input.txt', solve, console.log);