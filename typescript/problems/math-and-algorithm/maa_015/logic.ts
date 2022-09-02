export const logic = (arg: string) => {
  const [a, b] = arg.split(" ").map(Number);
  return gcd(a, b);
};

function gcd(a: number, b: number): number {
  if (b === 0) {
    return a;
  }

  return gcd(b, a % b);
}
