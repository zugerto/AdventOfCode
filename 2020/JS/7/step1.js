"use strict";

var files = require('../util/files');

const containsShinyGold = (map, color) => {
    if ('shiny gold' === color){
        return true;
    }
    if (!map.has(color)){
        return false;
    }

    const colorsInBag = map.get(color);
    for (const colorInBag of colorsInBag) {
        if (containsShinyGold(map, colorInBag)){
            return true;
        }
        
    }

    return false;
}

const solve = (input) => { 
    var lines = input.trim().split('\n')

    const map = new Map();
    
    for (const line of lines){
        const [bag, bags] = line.split(' bags contain ');
        
        bags.split(', ').map(b => {
            const {groups} = b.match(/((?<number>\d+) )?((?<color>.*) bags?\.?)/);
            if (!map.has(bag)){
                map.set(bag, [groups.color]);
            } else {
                map.set(bag, [... map.get(bag), groups.color]);
            }

        })
    }

    // The map now contains:  
    // 'light red' => [ 'bright white', 'muted yellow' ]
    // 'muted yellow' => [ 'shiny gold', 'faded blue' ], etc..                     
    
    let total = 0;
    for (const color of map.keys()){
        if (containsShinyGold(map, color)){
            total++;
        }
    }

    return total - 1;
}

console.assert(solve(
    'light red bags contain 1 bright white bag, 2 muted yellow bags.\n' +
    'dark orange bags contain 3 bright white bags, 4 muted yellow bags.\n' +
    'bright white bags contain 1 shiny gold bag.\n' +
    'muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\n' +
    'shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\n' +
    'dark olive bags contain 3 faded blue bags, 4 dotted black bags.\n' +
    'vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\n' +
    'faded blue bags contain no other bags.\n' +
    'dotted black bags contain no other bags.')==4)

files.processWithResult('input.txt', solve, console.log);