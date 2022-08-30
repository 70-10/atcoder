import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3\n30 50 70", 50],
  ["10\n1 2 3 4 5 6 7 8 9 10", 55],
  ["5\n60 60 60 60 60", 0],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
