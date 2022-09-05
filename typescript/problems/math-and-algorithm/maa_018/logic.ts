const TOAL = 500;

export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);

  let [a, b, c, d] = [0, 0, 0, 0];

  nums.forEach((num) => {
    switch (num) {
      case 100:
        a++;
        break;
      case 200:
        b++;
        break;
      case 300:
        c++;
        break;
      case 400:
        d++;
        break;
      default:
        break;
    }
  });

  return a * d + b * c;
};
