"use strict";

var files = require('../util/files');

const validateFourDigitsWithinLimit = (item, low, high) => {
    if (/\d{4}/.test(item)){
        var year = parseInt(item);
        if (year >= low && year <= high){
            return true;
        }
    } else {
        return false;
    }
}

const itemValue = (passportItems, field) => {
    return passportItems.filter(item => item.startsWith(field))[0].split(':')[1];
}
    
const isValidPassport = (passportItems) => {
    var itemTypes = passportItems.map(item => item.split(':')[0]);
    if (!['ecl', 'pid', 'eyr', 'hcl', 'byr', 'iyr', 'hgt'].every(v => itemTypes.includes(v))){
        return false;
    }

    //byr (Birth Year) - four digits; at least 1920 and at most 2002.
    var byr = itemValue(passportItems, 'byr');
    if (!validateFourDigitsWithinLimit(byr, 1920, 2002)){
        return false;
    }
    
    //iyr (Issue Year) - four digits; at least 2010 and at most 2020.
    var iyr = itemValue(passportItems, 'iyr');
    if (!validateFourDigitsWithinLimit(iyr, 2010, 2020)){
        return false;
    }

    //eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
    var eyr = itemValue(passportItems, 'eyr');
    if (!validateFourDigitsWithinLimit(eyr, 2020, 2030)){
        return false;
    }

    //hgt (Height) - a number followed by either cm or in:
    // If cm, the number must be at least 150 and at most 193.
    // If in, the number must be at least 59 and at most 76.
    var hgt = itemValue(passportItems, 'hgt');
    if (/\d+(in)/.test(hgt)){
        var height = parseInt(hgt.split('in')[0]);
        if (height < 59 || height > 76){
            return false;
        }
    } else if (/\d+(cm)/.test(hgt)){
        var height = parseInt(hgt.split('cm')[0]);
        if (height < 150 || height > 193){
            return false;
        }
    } else {
        return false;
    }   

    //hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
    var hcl = itemValue(passportItems, 'hcl');
    if (!/#(\d|[a-f]){6}/.test(hcl)){
        return false;
    } 

    //ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    var ecl = itemValue(passportItems, 'ecl');
    if (!/(amb|blu|brn|gry|grn|hzl|oth)/.test(ecl)){
        return false;
    } 

    //pid (Passport ID) - a nine-digit number, including leading zeroes.
    var pid = itemValue(passportItems, 'pid');
    if (!/\d{9}/.test(pid)){
        return false;
    } 
    
    return true;
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

//Invalid passports
console.assert(isValidPassport(['eyr:1972', 'cid:100',
'hcl:#18171d', 'ecl:amb', 'hgt:170', 'pid:186cm', 'iyr:2018', 'byr:1926']) == false)
console.assert(isValidPassport(['iyr:2019',
'hcl:#602927', 'eyr:1967', 'hgt:170cm',
'ecl:grn', 'pid:012533040', 'byr:1946']) == false)
console.assert(isValidPassport(['hcl:dab227', 'iyr:2012',
'ecl:brn', 'hgt:182cm', 'pid:021572410', 'eyr:2020', 'byr:1992', 'cid:277']) == false)
console.assert(isValidPassport(['hgt:59cm', 'ecl:zzz',
'eyr:2038', 'hcl:74454a', 'iyr:2023',
'pid:3556412378', 'byr:2007']) == false)

//Valid passports
console.assert(isValidPassport(['pid:087499704', 'hgt:74in', 'ecl:grn', 'iyr:2012', 'eyr:2030', 'byr:1980',
'hcl:#623a2f']) == true)
console.assert(isValidPassport(['eyr:2029', 'ecl:blu', 'cid:129', 'byr:1989',
'iyr:2014', 'pid:896056539', 'hcl:#a97842', 'hgt:165cm']) == true)
console.assert(isValidPassport(['hcl:#888785',
'hgt:164cm', 'byr:2001', 'iyr:2015', 'cid:88',
'pid:545766238', 'ecl:hzl',
'eyr:2022']) == true)
console.assert(isValidPassport(['iyr:2010', 'hgt:158cm', 'hcl:#b6652a', 'ecl:blu', 'byr:1944', 'eyr:2021', 'pid:093154719']) == true)

files.processLinesWithResult('input.txt', solve, console.log);