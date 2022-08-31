export const logic = (arg: string) => {
  const [N, a, b] = arg.split(" ").map(Number);
  let count = 0;

  for (let i = 1; i <= N; i++) {
    if (i % a === 0 || i % b === 0) {
      count++;
    }
  }
  return count;
};
