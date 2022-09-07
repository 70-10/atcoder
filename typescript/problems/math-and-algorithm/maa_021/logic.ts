export const logic = (arg: string) => {
  const [n, r] = arg.split(" ").map(Number);
  return factorial(n) / (factorial(r) * factorial(n - r));
};

function factorial(num: number): number {
  if (num < 2) {
    return 1;
  }

  return num * factorial(num - 1);
}
