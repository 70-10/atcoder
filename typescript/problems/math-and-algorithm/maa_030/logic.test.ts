import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["3 8\n3 30\n4 50\n5 60", 90],
  [
    "5 5\n1 1000000000\n1 1000000000\n1 1000000000\n1 1000000000\n1 1000000000",
    5000000000,
  ],
  ["6 15\n6 5\n5 6\n6 4\n6 6\n3 5\n7 2", 17],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
