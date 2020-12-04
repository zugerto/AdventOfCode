"use strict";

var files = require('../util/files');

let containsAll = (arr, target) => target.every(v => arr.includes(v));

const isValidPassport = (passportItems) => {
    var itemTypes = passportItems.map(item => item.split(':')[0]);
    if (containsAll(itemTypes, ['ecl', 'pid', 'eyr', 'hcl', 'byr', 'iyr', 'hgt'])){
        return true;
    }
    return false;
}

const solve = (rows) => { 
    var numberOfValidPassports = 0;
    var lineNumber = 0;
    var passportItems = [];
    
    while (lineNumber < rows.length){
        const matchedPassportItems = rows[lineNumber].match(/(\S+)\:(\S+)/g); 

        if (matchedPassportItems == null){
            if (isValidPassport(passportItems)){
                numberOfValidPassports++;
            }
            passportItems = [];
        } else {
            passportItems = [...passportItems, ...matchedPassportItems];
        }

        lineNumber++;
    }
    
    if (isValidPassport(passportItems)){
        numberOfValidPassports++;
    }

    return numberOfValidPassports;
}

console.assert(isValidPassport([ 'ecl:gry', 'pid:860033327', 'eyr:2020', 'hcl:#fffffd',
'byr:1937', 'iyr:2017', 'cid:147', 'hgt:183cm' ]) == true)

console.assert(isValidPassport([ 'iyr:2013', 'ecl:amb', 'cid:350', 'eyr:2023', 'pid:028048884', 
 'hcl:#cfa07d', 'byr:1929' ]) == false);

console.assert(isValidPassport([ 'hcl:#ae17e1', 'iyr:2013', 
'eyr:2024', 
'ecl:brn', 'pid:760753108', 'byr:1931', 
'hgt:179cm' ]) == true);

console.assert(isValidPassport([ 'hcl:#cfa07d', 'eyr:2025', 'pid:166559648', 
 'iyr:2011', 'ecl:brn', 'hgt:59in' ]) == false);

console.assert(solve([
    'ecl:gry pid:860033327 eyr:2020 hcl:#fffffd',
    'byr:1937 iyr:2017 cid:147 hgt:183cm',
    '',
    'iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884',
    'hcl:#cfa07d byr:1929',
    '',
    'hcl:#cfa07d eyr:2025 pid:166559648',
    'iyr:2011 ecl:brn hgt:59in',
    '',
    'hcl:#ae17e1 iyr:2013',
    'eyr:2024',
    'ecl:brn pid:760753108 byr:1931',
    'hgt:179cm',])==2)

files.processLinesWithResult('input.txt', solve, console.log);