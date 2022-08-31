export const logic = (arg: string) => {
  const n = Number(arg);
  const nums: number[] = [];

  for (let i = 1; i * i <= n; i++) {
    if (n % i === 0) {
      nums.push(i);
      if (n / i !== i) {
        nums.push(n / i);
      }
    }
  }
  return nums.sort((a, b) => a - b).join("\n");
};
