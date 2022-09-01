import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["10", "2 5"],
  ["36", "2 2 3 3"],
  ["460", "2 2 5 23"],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
