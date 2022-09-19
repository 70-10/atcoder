import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["7 3\n1 2 3 4 5 6 7", "Yes"],
  ["7 9\n1 2 3 4 5 6 7", "No"],
  ["7 1\n2 3 4 5 6 7 8", "No"],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
