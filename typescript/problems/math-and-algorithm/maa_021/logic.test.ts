import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["6 2", 15],
  ["19 4", 3876],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
