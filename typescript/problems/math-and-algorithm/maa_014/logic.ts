export const logic = (arg: string) => {
  let n = Number(arg);
  const result: number[] = [];

  for (let i = 2; i * i <= n; i++) {
    while (n % i === 0) {
      result.push(i);
      n = n / i;
    }
  }

  if (n !== 1) {
    result.push(n);
  }

  return result.join(" ");
};
