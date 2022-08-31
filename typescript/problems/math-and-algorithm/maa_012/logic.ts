export const logic = (arg: string) => {
  const n = Number(arg);
  return isPrime(n) ? "Yes" : "No";
};

function isPrime(num: number) {
  if (num <= 1) {
    return false;
  }

  for (let i = 2; i * i < num; i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}
