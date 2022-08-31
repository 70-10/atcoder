export const logic = (arg: string) => {
  const [n, s] = arg.split(" ").map(Number);
  let count = 0;

  for (let i = 1; i <= n; i++) {
    for (let j = 1; j <= n; j++) {
      if (i + j <= s) {
        count++;
      }
    }
  }
  return count;
};
