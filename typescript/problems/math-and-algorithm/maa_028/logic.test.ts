import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["4\n10 30 40 20", 30],
  ["2\n10 10", 0],
  ["6\n30 10 60 10 60 50", 40],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
