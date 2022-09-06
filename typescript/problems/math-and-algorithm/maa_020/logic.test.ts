import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["5\n100 150 200 250 300", 1],
  ["13\n243 156 104 280 142 286 196 132 128 195 265 300 130", 4],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
