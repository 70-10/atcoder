import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([["5\n100 300 400 400 200", 3]])(
  "logic(%s) -> %s",
  (arg, expected) => {
    expect(logic(arg)).toBe(expected);
  }
);
