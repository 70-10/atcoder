export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);

  const result: number[] = [];

  for (let i = 0; i < nums.length; i++) {
    if (i === 0) {
      result[i] = 0;
    }

    if (i === 1) {
      result[i] = Math.abs(nums[i - 1] - nums[i]);
    }

    if (i >= 2) {
      const v1 = result[i - 1] + Math.abs(nums[i - 1] - nums[i]);
      const v2 = result[i - 2] + Math.abs(nums[i - 2] - nums[i]);
      result[i] = Math.min(v1, v2);
    }
  }

  return result[nums.length - 1];
};
