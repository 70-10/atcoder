import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["2 8 8", 128],
  ["7 7 25", 1225],
  ["100 100 100", 1000000],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
