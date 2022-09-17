import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["5\n2 5 3 3 1", 8]])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
