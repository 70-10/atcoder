export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);
  return mergeSort(nums).join(" ");
};

function mergeSort(nums: number[]): number[] {
  if (nums.length === 1) {
    return nums;
  }

  const i = Math.floor(nums.length / 2);
  const a = mergeSort(nums.slice(0, i));
  const b = mergeSort(nums.slice(i, nums.length));

  let c1 = 0;
  let c2 = 0;
  const result = [];

  while (c1 < a.length || c2 < b.length) {
    if (c1 === a.length || a[c1] > b[c2]) {
      result.push(b[c2]);
      c2++;
      continue;
    }

    result.push(a[c1]);
    c1++;
  }

  return result;
}
