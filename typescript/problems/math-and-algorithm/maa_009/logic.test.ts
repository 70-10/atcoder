import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3 11\n2 5 9", "Yes"],
  ["4 11\n3 1 4 5", "No"],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
