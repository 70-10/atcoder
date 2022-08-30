export const logic = (arg: string) => {
  return arg
    .split(" ")
    .map(Number)
    .reduce((sum, num) => sum * num, 1);
};
