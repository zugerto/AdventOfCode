"use strict";

var files = require('../util/files');

const getSeatId = (boardingpass) => {
    var row = parseInt(boardingpass
                .substring(0,7)
                .replace(/F/g, '0')
                .replace(/B/g, '1'), 2);
    var column = parseInt(boardingpass
                    .substring(7)
                    .replace(/L/g, '0')
                    .replace(/R/g, '1'), 2);

    var seatId = row * 8 + column;    
    console.log(boardingpass, ': row ', row, ', column ', column, ' seat ID ', seatId)

    return seatId;
}

const solve = (boardingpasses) => { 
    return Math.max(...boardingpasses.map(boardingpass => getSeatId(boardingpass)));
}

//BFFFBBFRRR: row 70, column 7, seat ID 567.
console.assert(getSeatId('BFFFBBFRRR') == 567)
//FFFBBBFRRR: row 14, column 7, seat ID 119.
console.assert(getSeatId('FFFBBBFRRR') == 119)
//BBFFBBFRLL: row 102, column 4, seat ID 820.
console.assert(getSeatId('BBFFBBFRLL') == 820)

console.assert(solve([
    'BFFFBBFRRR',
    'FFFBBBFRRR',
    'BBFFBBFRLL',
   ])==820)

files.processLinesWithResult('input.txt', solve, console.log);