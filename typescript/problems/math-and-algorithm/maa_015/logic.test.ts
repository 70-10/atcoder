import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["33 88", 11],
  ["88 33", 11],
  ["123 777", 3],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
