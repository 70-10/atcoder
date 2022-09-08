export const logic = (arg: string) => {
  const n = Number(arg.split("\n")[0]);
  let result = 0;

  for (let i = 1; i <= n; i++) {
    const [p, q] = arg.split("\n")[i].split(" ").map(Number);
    result += q / p;
  }
  return result;
};
