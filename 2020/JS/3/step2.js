"use strict";

const { count } = require('console');
var files = require('../util/files');

const countTree = (rows, right, down) => { 
    var numTrees = 0, x = 0;

    for(var y = down; y < rows.length; y += down){
        if (rows[y][(x += right)%(rows[y].length)] == '#') {
            numTrees++;
        }
    }

    return numTrees;
}

const solve = (rows) => {
    return countTree(rows,1,1) 
        * countTree(rows,3,1) 
        * countTree(rows,5,1)  
        * countTree(rows,7,1) 
        * countTree(rows,1,2);
}

const map = [
    '..##.........##.........##.........##.........##.........##.......',
    '#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..',
    '.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.',
    '..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#',
    '.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.',
    '..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....',
    '.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#',
    '.#........#.#........#.#........#.#........#.#........#.#........#',
    '#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...',
    '#...##....##...##....##...##....##...##....##...##....##...##....#',
    '.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#'];
console.assert(countTree(map,1,1)==2)
console.assert(countTree(map,3,1)==7)
console.assert(countTree(map,5,1)==3)  
console.assert(countTree(map,7,1)==4)
console.assert(countTree(map,1,2)==2)        
console.assert(solve(map)==336)

files.processLinesWithResult('input.txt', solve, console.log);