export const logic = (arg: string) => {
  const n = Number(arg);

  const result: number[] = [];

  for (let i = 0; i <= n; i++) {
    if (i <= 1) {
      result[i] = 1;
    } else {
      result[i] = result[i - 1] + result[i - 2];
    }
  }

  return result[n];
};
