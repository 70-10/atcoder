export const logic = (arg: string) => {
  const [a, b, c] = arg.split(" ").map(Number);
  return a + b + c;
};
