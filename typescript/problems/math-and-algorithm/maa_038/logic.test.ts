import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  [
    "10 5\n8 6 9 1 2 1 10 100 1000 10000\n2 3\n1 4\n3 9\n6 8\n1 10",
    "15\n24\n1123\n111\n11137",
  ],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
