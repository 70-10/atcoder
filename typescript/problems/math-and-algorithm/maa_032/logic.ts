export const logic = (arg: string) => {
  const [n, x] = arg.split("\n")[0].split(" ").map(Number);
  const nums = arg.split("\n")[1].split(" ").map(Number);

  return nums.map((num) => num === x).includes(true) ? "Yes" : "No";
};
