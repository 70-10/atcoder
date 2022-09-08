export const logic = (arg: string) => {
  const n = Number(arg.split("\n")[0]);
  const blue = arg.split("\n")[1].split(" ").map(Number);
  const red = arg.split("\n")[2].split(" ").map(Number);

  return sum(blue) / n + sum(red) / n;
};

function sum(nums: number[]): number {
  return nums.reduce((total, num) => total + num, 0);
}
