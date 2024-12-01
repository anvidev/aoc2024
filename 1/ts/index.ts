const data = await Deno.readTextFile("../data.txt");
const { left, right } = splitAndSortData(data);

console.log(left, right);

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

//function calculateDistance(left: number[], right: number[]): number {}
