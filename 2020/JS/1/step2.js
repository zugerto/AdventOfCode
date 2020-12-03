"use strict";

var files = require('../util/files');

function solve(entries){
    for (var x of entries) {
        for (var y of entries) {
            for (var z of entries) {
                if (parseInt(x) + parseInt(y) + parseInt(z) == 2020){
                    return parseInt(x) * parseInt(y) * parseInt(z);
                }
            }
        }
    }
}

console.assert(solve([1721,979,366,299,675,1456])==241861950)

files.processLinesWithResult('input.txt', solve, console.log);