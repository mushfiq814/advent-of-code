import { readFileSync } from 'fs';

enum Hand {
  ROCK = 1,
  PAPER = 2,
  SCISSOR = 3,
};

enum Result {
  WIN = 6,
  DRAW = 3,
  LOSE = 0,
};

type Input = 'A' | 'B' | 'C' | 'X' | 'Y' | 'Z';

const inputFile = 'input.txt'
const input = readFileSync(inputFile, 'utf8');
const parsedRounds = input
  .split('\n')
  .map(x => x.replace(
    new RegExp('^[\t\s ]+', 'g'),
    ''
  ))
  .filter(x => x !== '')
  .map(x => x.split(' '))

const partOne = (parsedRounds: string[][]) => {
  const handMapper = (x: Input) => {
    if (x === 'A' || x === 'X') return Hand.ROCK;
    if (x === 'B' || x === 'Y') return Hand.PAPER;
    if (x === 'C' || x === 'Z') return Hand.SCISSOR;
    return Hand.ROCK;
  }

  const rounds = parsedRounds.map(round => ({
    opponent: handMapper(round[0] as Input),
    own: handMapper(round[1] as Input),
  }));

  const getGameResult = (opponent: Hand, own: Hand): Result => {
    if (opponent === own) return Result.DRAW;

    if (
      (opponent === Hand.ROCK && own === Hand.PAPER)
      || (opponent === Hand.PAPER && own === Hand.SCISSOR)
      || (opponent === Hand.SCISSOR && own === Hand.ROCK)
    ) return Result.WIN;

    return Result.LOSE;
  }

  return rounds.reduce((score, round) =>
    score += round.own + getGameResult(round.opponent, round.own),
    0
  );
}

const partTwo = (parsedRounds: string[][]) => {
  const handMapper = (x: Input) => {
    if (x === 'A') return Hand.ROCK;
    if (x === 'B') return Hand.PAPER;
    if (x === 'C') return Hand.SCISSOR;
    if (x === 'X') return Result.LOSE;
    if (x === 'Y') return Result.DRAW;
    if (x === 'Z') return Result.WIN;
    return Hand.ROCK;
  }

  const rounds = parsedRounds.map(round => ({
      opponent: handMapper(round[0] as Input),
      result: handMapper(round[1] as Input),
    }));

  const getGameResult = (opponent: Hand, result: Result): Hand => {
    if (result === Result.DRAW) return opponent;
    if (result === Result.WIN) {
      if (opponent === Hand.ROCK) return Hand.PAPER;
      if (opponent === Hand.PAPER) return Hand.SCISSOR;
      if (opponent === Hand.SCISSOR) return Hand.ROCK;
    }
    if (opponent === Hand.ROCK) return Hand.SCISSOR;
    if (opponent === Hand.PAPER) return Hand.ROCK;
    return Hand.PAPER;
  }

  return rounds.reduce((score, round) =>
    score += round.result + getGameResult(
      round.opponent as Hand,
      round.result as Result
    ), 0);
}

console.log(partOne(parsedRounds));
console.log(partTwo(parsedRounds));
