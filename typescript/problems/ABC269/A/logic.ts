export const logic = (arg: string) => {
  const [a, b, c, d] = arg.split(" ").map(Number);
  return `${(a + b) * (c - d)}\nTakahashi`;
};
