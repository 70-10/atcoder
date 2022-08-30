import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["100", 203],
  ["27", 57],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
