'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
    inputString += inputStdin;
});

process.stdin.on('end', _ => {
    inputString = inputString.trim().split('\n').map(string => {
        return string.trim();
    });

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the vowelsAndConsonants function.
 * Print your output using 'console.log()'.
 */
function vowelsAndConsonants(s) {
    let vowels = [], consonants = [];
    for (let c of s) {
        if (['a', 'e', 'i', 'o', 'u'].indexOf(c, 0) !== -1)
            vowels.push(c);
        else
            consonants.push(c);
    }
    vowels.forEach((v) => console.log(v));
    consonants.forEach((c) => console.log(c));
}


function main() {
    const s = readLine();

    vowelsAndConsonants(s);
}