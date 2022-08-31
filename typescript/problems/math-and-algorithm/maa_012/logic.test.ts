import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["53", "Yes"],
  ["77", "No"],
  ["472249589291", "Yes"],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
