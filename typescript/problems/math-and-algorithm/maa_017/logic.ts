export const logic = (arg: string) => {
  const nums = arg
    .split("\n")[1]
    .split(" ")
    .map((i) => BigInt(i));
  return Number(nums.reduce((g, num) => (g * num) / gcd(g, num), nums[0]));
};

function gcd(a: bigint, b: bigint): bigint {
  if (b === BigInt(0)) {
    return a;
  }

  return gcd(b, a % b);
}
