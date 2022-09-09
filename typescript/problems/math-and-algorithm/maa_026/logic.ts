export const logic = (arg: string) => {
  const n = Number(arg);
  return (
    Math.ceil(
      [...Array(n)]
        .map((_, i) => i)
        .reduce((total, num) => total + n / (n - num), 0) *
        10 ** 12
    ) /
    10 ** 12
  );
};
