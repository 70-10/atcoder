import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["5\n3 1 4 1 5\n9 2 6 5 3", 21.333333333333]])(
  "logic(%s) -> %s",
  (arg, expected) => {
    expect(logic(arg)).toBe(expected);
  }
);
