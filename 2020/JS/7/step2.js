"use strict";

var files = require('../util/files');

const sum = (map, bag) => {
    if (bag.number == 0){
        return 0;
    }

    let sumOfBags = 1;
    const bagsInBag = map.get(bag.color);
    for (const bagInBags of bagsInBag) {
        sumOfBags += bagInBags.number * sum(map, bagInBags)
    }

    return sumOfBags;
}

const solve = (input) => { 
    var lines = input.trim().split('\n')

    const map = new Map();
    
    for (const line of lines){
        const [bag, bags] = line.split(' bags contain ');
        
        bags.split(', ').map(b => {
            const {groups} = b.match(/((?<number>\d+) )?((?<color>.*) bags?\.?)/);
            
            //No other bags is 0
            if (!groups.number){
                groups.number = 0;
            } 

            if (!map.has(bag)){
                map.set(bag, [groups]);
            } else {
                map.set(bag, [... map.get(bag), groups]);
            }

        })
    }

    // The map now contains:  
    //  'shiny magenta' => [{ number: '3', color: 'faded coral' },
    //                     { number: '5', color: 'posh gray' }], etc..
                    
    return sum(map, {number: 1, color: 'shiny gold'}) - 1;
}

console.assert(solve(
    'shiny gold bags contain 2 dark red bags.\n' +
    'dark red bags contain 2 dark orange bags.\n' +
    'dark orange bags contain 2 dark yellow bags.\n' +
    'dark yellow bags contain 2 dark green bags.\n' +
    'dark green bags contain 2 dark blue bags.\n' +
    'dark blue bags contain 2 dark violet bags.\n' +
    'dark violet bags contain no other bags.')==126)

files.processWithResult('input.txt', solve, console.log);