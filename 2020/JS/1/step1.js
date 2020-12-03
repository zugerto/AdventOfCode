"use strict";

var files = require('../util/files');

function solve(entries){
    for (var x of entries) {
        for (var y of entries) {
            if (parseInt(x) + parseInt(y) == 2020){
                return parseInt(x) * parseInt(y);
            }
        }
    }
}

console.assert(solve([1721,979,366,299,675,1456])==514579)

files.processLinesWithResult('input.txt', solve, console.log);