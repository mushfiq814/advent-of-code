import { readFileSync } from "fs";

// const inputFile = "input-short.txt";
const inputFile = "input.txt";
const input = readFileSync(inputFile, "utf8");

const pairs = input
  .split("\n")
  .filter((x) => x != "")
  .map((x) => x.split(","));

const partOne = (pairs: string[][]) => {
  const parseRange = (range: string) =>
    range.split("-").map((num) => parseInt(num));

  return pairs
    .map((pair) => {
      const [start1, end1] = parseRange(pair[0]);
      const [start2, end2] = parseRange(pair[1]);

      return (
        (start2 >= start1 && end2 <= end1) || (start1 >= start2 && end1 <= end2)
      );
    })
    .reduce((count, isOverlap) => (count += isOverlap ? 1 : 0), 0);
};

console.log(partOne(pairs));
