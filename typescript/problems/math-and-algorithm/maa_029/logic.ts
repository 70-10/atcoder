export const logic = (arg: string) => {
  const n = Number(arg);
  const patterns: number[] = [];

  for (let i = 0; i <= n; i++) {
    if (i <= 1) {
      patterns[i] = 1;
      continue;
    }

    patterns[i] = patterns[i - 1] + patterns[i - 2];
  }

  return patterns[patterns.length - 1];
};
