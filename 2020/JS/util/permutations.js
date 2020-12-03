"use strict";

//implements Heaps algorithm for calculating permutations, https://en.wikipedia.org/wiki/Heap%27s_algorithm

exports.heapsAlgorithm = function(array) {
    const result = []
    const addToResult = (item) => {
     result.push(item.slice());
   }
   generate(array.length, array, addToResult);
   
   return result;
}

// Generate the permutation for a given n (amount of elements) and a given array
function generate(n, arr, output) {
 
  // If only 1 element, just output the array
  if (n == 1) {
    output(arr);
  }

  for (var i = 0; i < n; i+= 1) {
    generate(n - 1, arr, output);

    // If n is even
    if (n % 2 == 0) {
      swap(arr, i, n - 1);
    } else {
      swap(arr, 0, n - 1);
    }
  }
}

function swap(arr, idxA, idxB) {
  var tmp = arr[idxA];
  arr[idxA] = arr[idxB];
  arr[idxB] = tmp;
}