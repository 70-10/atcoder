export const logic = (arg: string) => {
  const nums = arg.split("\n")[1].split(" ").map(Number);

  let [one, two, three] = [0, 0, 0, 0];

  nums.forEach((num) => {
    switch (num) {
      case 1:
        one++;
        break;
      case 2:
        two++;
        break;
      case 3:
        three++;
        break;
      default:
        break;
    }
  });

  return floorSum(one - 1) + floorSum(two - 1) + floorSum(three - 1);
};

function floorSum(num: number) {
  return ((num + 1) * num) / 2;
}
