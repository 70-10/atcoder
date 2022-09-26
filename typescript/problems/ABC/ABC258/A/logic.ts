export const logic = (num: number): string => {
  return `${21 + Math.trunc(num / 60)}:${("0" + (num % 60)).slice(-2)}`;
};
