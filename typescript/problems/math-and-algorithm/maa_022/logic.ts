export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);
  const n = 99999;
  const counts = Array(n).fill(0);
  let count = 0;

  nums.forEach((num) => {
    counts[num - 1]++;
  });

  for (let i = 0; i < (n - 1) / 2; i++) {
    count += counts[i] * counts[n - 1 - i];
  }

  const m = Math.floor(n / 2);
  count += (counts[m] * (counts[m] - 1)) / 2;
  return count;
};
