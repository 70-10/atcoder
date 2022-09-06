import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["6\n1 3 2 1 1 2", 4]])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
