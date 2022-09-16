export const logic = (arg: string) => {
  const args = arg.split("\n");
  const [N, W] = args[0].split(" ").map(Number);
  const nums = [...args].map((s) => s.split(" ").map(Number));
  nums.shift();

  const dp: number[][] = [...Array(N + 1)].map(() => Array(W + 1).fill(0));

  for (let i = 1; i <= N; i++) {
    const [w, v] = nums[i - 1];

    for (let j = 0; j <= W; j++) {
      const a = dp[i - 1][j];
      const b = j - w >= 0 ? dp[i - 1][j - w] + v : 0;
      dp[i][j] = Math.max(a, b);
    }
  }

  return Array.from(new Set(dp.flat())).sort((a, b) => b - a)[0];
};
