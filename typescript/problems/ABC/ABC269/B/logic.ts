export const logic = (arg: string) => {
  const list = arg.split("\n");
  const result = { a: 0, b: 0, c: 0, d: 0 };

  for (let i = 0; i < list.length; i++) {
    const s = list[i];
    if (!s.includes("#")) {
      continue;
    }

    result.a = result.a === 0 ? i + 1 : result.a;
    result.b = i + 1;

    if (result.c !== 0 || result.d !== 0) {
      continue;
    }

    const arr = [...s];
    for (let j = 0; j < arr.length; j++) {
      if (arr[j] !== "#") {
        continue;
      }
      result.c = result.c === 0 ? j + 1 : result.c;
      result.d = j + 1;
    }
  }

  const { a, b, c, d } = result;
  return `${a} ${b}\n${c} ${d}`;
};
