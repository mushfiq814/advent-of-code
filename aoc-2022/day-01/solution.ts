import { readFileSync } from 'fs';

const inputFile = 'input.txt'
const input = readFileSync(inputFile, 'utf8');

const caloriesPerPerson =
  input
    .replace(new RegExp(' ', 'g'), '')
    .split('\n\n')
    .map(x => x
      .split('\n')
      .filter(x => x !== '')
      .map(x => parseInt(x))
    );

const totalCaloriesPerPerson =
  caloriesPerPerson.map(cals => cals.reduce((a, b) => a + b, 0));

const maxCalories = Math.max(...totalCaloriesPerPerson);

const highestCaloryElfIndex = totalCaloriesPerPerson.indexOf(maxCalories);

console.log(`Elf #${highestCaloryElfIndex + 1} is carrying ${maxCalories}`);
