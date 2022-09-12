import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3\n3 1 2", "1 2 3"],
  [
    "10\n658 299 47 507 122 969 449 68 513 800",
    "47 68 122 299 449 507 513 658 800 969",
  ],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
