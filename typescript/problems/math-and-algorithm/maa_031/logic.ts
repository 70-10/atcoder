export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);

  const dp1: number[] = [];
  const dp2: number[] = [];

  for (let i = 0; i <= nums.length; i++) {
    switch (i) {
      case 0: {
        dp1[i] = 0;
        dp2[i] = 0;
        break;
      }
      case 1: {
        dp1[i] = nums[i - 1];
        dp2[i] = 0;
        break;
      }
      default: {
        dp1[i] = dp2[i - 1] + nums[i - 1];
        dp2[i] = Math.max(dp1[i - 1], dp2[i - 1]);
      }
    }
  }

  return Math.max(dp1[dp1.length - 1], dp2[dp2.length - 1]);
};
