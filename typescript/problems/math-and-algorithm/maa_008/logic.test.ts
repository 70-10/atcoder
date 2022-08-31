import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3 4", 6],
  ["869 120", 7140],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
