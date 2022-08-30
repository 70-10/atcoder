import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["10 20 50", 80],
  ["1 1 1", 3],
  ["100 100 100", 300],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
