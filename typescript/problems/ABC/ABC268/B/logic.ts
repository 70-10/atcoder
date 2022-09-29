export const logic = (arg: string) => {
  const [S, T] = arg.split("\n");

  if (S.length > T.length) {
    return "No";
  }

  return T.startsWith(S) ? "Yes" : "No";
};
