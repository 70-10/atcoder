import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["10", "2 3 5 7"],
  ["4", "2 3"],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
