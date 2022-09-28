import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["31 9 24 31 24", 3],
  ["0 0 0 0 0", 1],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
