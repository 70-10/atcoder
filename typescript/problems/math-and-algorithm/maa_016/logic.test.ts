import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3\n12 18 24", 6],
  ["4\n10 20 100 1000", 10],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
