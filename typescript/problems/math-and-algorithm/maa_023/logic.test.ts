import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["3\n1 2 3\n10 20 30", 22.0]])(
  "logic(%s) -> %s",
  (arg, expected) => {
    expect(logic(arg)).toBe(expected);
  }
);
