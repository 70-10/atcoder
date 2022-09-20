export const logic = (arg: string) => {
  const [N, Q] = arg.split("\n")[0].split(" ").map(Number);
  const counts = arg.split("\n")[1].split(" ").map(Number);

  const args = arg.split("\n");
  args.shift();
  args.shift();

  const { l, r } = args.reduce(
    (obj, a) => {
      const [x, y] = a.split(" ").map(Number);
      obj.l.push(x);
      obj.r.push(y);
      return obj;
    },
    { l: [], r: [] }
  );

  const accum = [...new Array(N + 1)]
    .map((_, i) => i)
    .reduce((arr, i) => {
      if (i === 0) {
        arr.push(0);
        return arr;
      }

      arr.push(counts[i - 1] + arr[i - 1]);
      return arr;
    }, []);

  const result = [...new Array(Q)].map((_, i) => {
    return accum[r[i]] - accum[l[i] - 1];
  });

  return result.join("\n");
};
