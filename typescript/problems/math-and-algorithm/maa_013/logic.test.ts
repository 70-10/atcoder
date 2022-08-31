import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  ["12", "1\n2\n3\n4\n6\n12"],
  ["827847039317", "1\n909859\n909863\n827847039317"],
])("logic(%s) -> %i", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
