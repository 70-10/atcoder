import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["2\n2 50\n4 100", 50.0]])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
