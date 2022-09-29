import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["atco\natcoder", "Yes"],
  ["code\natcoder", "No"],
  ["abc\nabc", "Yes"],
  ["aaaa\naa", "No"],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
