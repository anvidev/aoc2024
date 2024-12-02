const data = await Deno.readTextFile("../data.txt");
const { left, right } = splitAndSortData(data);
const distance = calculateDistance(left, right);
const similarityScore = calculateSimilarityScore(left, right);

console.log({ distance, similarityScore });

function splitAndSortData(str: string): { left: number[]; right: number[] } {
  const numbers = str.split("\n", 1000);
  const left: number[] = [];
  const right: number[] = [];

  for (const line of numbers) {
    const pair = line.trim().split("   ");

    if (pair.length != 2 || isNaN(Number(pair[0])) || isNaN(Number(pair[1]))) {
      throw new Error("invalid line format");
    }

    left.push(Number(pair[0]));
    right.push(Number(pair[1]));
  }

  function sortNumberAscending(a: number, b: number): number {
    if (a < b) {
      return -1;
    } else if (a == b) {
      return 1;
    } else {
      return 1;
    }
  }

  left.sort((a, b) => sortNumberAscending(a, b));
  right.sort((a, b) => sortNumberAscending(a, b));

  return { left, right };
}

function calculateDistance(left: number[], right: number[]): number {
  if (left.length != right.length) {
    throw new Error("left and right is not equal in size");
  }
  let total = 0;
  for (let i = 0; i < left.length; i++) {
    total += Math.abs(left[i] - right[i]);
  }
  return total;
}

function calculateSimilarityScore(left: number[], right: number[]): number {
  const freq: { [key: number]: number } = {};
  for (const number of right) {
    freq[number] = (freq[number] || 0) + 1;
  }
  let total = 0;
  for (const number of left) {
    total += number * (freq[number] || 0);
  }
  return total;
}
