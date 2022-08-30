import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["5\n3 1 4 1 5", 14],
  ["3\n10 20 50", 80],
  ["10\n1 2 3 4 5 6 7 8 9 10", 55],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
