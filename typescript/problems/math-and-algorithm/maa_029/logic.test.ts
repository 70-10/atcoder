import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["4", 5],
  ["5", 8],
  ["45", 1836311903],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
