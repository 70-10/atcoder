import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["6", 720],
  ["5", 120],
  ["4", 24],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
