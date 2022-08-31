import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["15 3 5", 7],
  ["1000000 11 13", 160839],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
