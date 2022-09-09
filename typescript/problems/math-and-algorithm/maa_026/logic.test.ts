import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["5", 11.416666666667]])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
