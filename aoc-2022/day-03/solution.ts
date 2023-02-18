import { readFileSync } from 'fs';

const inputFile = 'input.txt'
const input = readFileSync(inputFile, 'utf8');

const rucksacks = input
  .split("\n")
  .filter(x => x !== '')
  .map(str => {
    const middle = str.length / 2;
    return [
      str.substr(0, middle),
      str.substr(middle),
    ];
  });

const partOne = (rucksacks: string[][]) => {
  const findRecurringVal = (x: string[], y: string[]) => {
    for (let i of x) {
      for (let j of y) {
        if (i === j) return i;
      }
    }
    return null;
  }

  return rucksacks.map(rucksack => {
    const x = rucksack[0].split("").sort();
    const y = rucksack[1].split("").sort();

    const recurringVal = findRecurringVal(x, y);

    if (!recurringVal) return null;

    // A -> 65, Z -> 90, a -> 97, z -> 122
    const unicodeValue = recurringVal.charCodeAt(0);
    const priority = (unicodeValue >= 97 && unicodeValue <= 122)
      ? unicodeValue - 96
      : (unicodeValue >= 65 && unicodeValue <= 90)
        ? unicodeValue - 38
        : null;

    return priority;
  }).reduce((sum: number, curr) => sum += curr ?? 0, 0);
}

console.log(partOne(rucksacks));
