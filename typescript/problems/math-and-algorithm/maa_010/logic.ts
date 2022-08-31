export const logic = (arg: string) => {
  const n = Number(arg);

  return [...Array(n)].map((_, i) => i + 1).reduce((sum, num) => sum * num, 1);
};
