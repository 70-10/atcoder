import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["1 2 5 3", "6\nTakahashi"],
  ["10 -20 30 -40", "-700\nTakahashi"],
  ["100 100 100 -100", "40000\nTakahashi"],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
