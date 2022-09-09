export const logic = (arg: string) => {
  const n = Number(arg.split("\n")[0]);
  const a = arg.split("\n")[1].split(" ").map(Number);
  const b = arg.split("\n")[2].split(" ").map(Number);

  let count = 0;

  for (let i = 0; i < n; i++) {
    count += a[i] * (2 / 6) + b[i] * (4 / 6);
  }
  return Math.floor(count * 10 ** 12) / 10 ** 12;
};
