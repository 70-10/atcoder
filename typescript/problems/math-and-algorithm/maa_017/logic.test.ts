import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3\n12 18 14", 252],
  ["4\n10 20 100 1000", 1000],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
