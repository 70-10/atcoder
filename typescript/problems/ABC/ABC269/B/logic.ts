export const logic = (arg: string) => {
  const list = arg.split("\n");
  let a = 0;
  let b = 0;
  let c = 0;
  let d = 0;

  for (let i = 0; i < list.length; i++) {
    const s = list[i];
    if (s.includes("#")) {
      if (a === 0) {
        a = i + 1;
      }
      b = i + 1;
    }

    if (c === 0 && d === 0) {
      const arr = [...s];
      for (let j = 0; j < arr.length; j++) {
        if (arr[j] === "#") {
          if (c === 0) {
            c = j + 1;
          }
          d = j + 1;
        }
      }
    }
  }

  return `${a} ${b}\n${c} ${d}`;
};
