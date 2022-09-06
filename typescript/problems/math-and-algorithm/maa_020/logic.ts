export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);

  let count = 0;

  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      for (let k = j + 1; k < nums.length; k++) {
        for (let l = k + 1; l < nums.length; l++) {
          for (let m = l + 1; m < nums.length; m++) {
            if (nums[i] + nums[j] + nums[k] + nums[l] + nums[m] === 1000) {
              count++;
            }
          }
        }
      }
    }
  }
  return count;
};
