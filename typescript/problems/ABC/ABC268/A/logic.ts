export const logic = (arg: string) => {
  const nums = arg.split(" ").map(Number);

  return [...new Set(nums)].length;
};
