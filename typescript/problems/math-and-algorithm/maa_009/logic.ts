export const logic = (arg: string) => {
  const [N, S] = arg.split("\n")[0].split(" ").map(Number);
  const nums = arg.split("\n")[1].split(" ").map(Number);

  const dp: boolean[][] = [...Array(N + 1)].map(() => [
    ...Array(S + 1).fill(false),
  ]);

  for (let i = 1; i <= N; i++) {
    for (let j = 0; j <= S; j++) {
      const numIndex = i - 1;
      const num = nums[numIndex];

      if (num === j || dp[numIndex][j] || dp[numIndex][j - num]) {
        dp[i][j] = true;
      }
    }
  }

  return dp[N][S] ? "Yes" : "No";
};
