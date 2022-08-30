export const logic = (arg: string) => {
  return (
    arg
      .split("\n")[1]
      .split(" ")
      .map(Number)
      .reduce((sum, num) => sum + num, 0) % 100
  );
};
