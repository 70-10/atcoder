import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  [1, 6],
  [100, 105],
])("logic(%i) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
