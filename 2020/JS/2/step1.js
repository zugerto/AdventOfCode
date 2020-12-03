"use strict";

var files = require('../util/files');

 const isValidPassword = (entry) => { 
    const matchedGroups = entry.match(/(\d+)\-(\d+)\s(\w)\:\s(\w+)/); 

    const low = parseInt(matchedGroups[1]);
    const high = parseInt(matchedGroups[2]);
    const needle = matchedGroups[3];
    const haystack = matchedGroups[4].split('');

    const numberOfMatchesInPassword = haystack.filter(x => x == needle).length;
    
    if (numberOfMatchesInPassword >= low && numberOfMatchesInPassword <= high){
        return 1;
    } else {
        return 0;
    }
 } 

function solve(entries){
    return entries.map(entry => {
       return isValidPassword(entry);
    }).reduce((a,b) => a + b, 0);
}

console.assert(solve(['1-3 a: abcde','1-3 b: cdefg','2-9 c: ccccccccc'])==2)

files.processLinesWithResult('input.txt', solve, console.log);