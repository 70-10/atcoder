import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["6\n40000 50000 20000 80000 50000 30000", 2]])(
  "logic(%s) -> %s",
  (arg, expected) => {
    expect(logic(arg)).toBe(expected);
  }
);
