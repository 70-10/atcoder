import { expect, test } from "vitest";
import { logic } from "./logic";

test.each([
  [
    "..........\n..........\n..........\n..........\n...######.\n...######.\n...######.\n...######.\n..........\n..........",
    "5 8\n4 9",
  ],
  [
    "..........\n..#.......\n..........\n..........\n..........\n..........\n..........\n..........\n..........\n..........",
    "2 2\n3 3",
  ],
  [
    "##########\n##########\n##########\n##########\n##########\n##########\n##########\n##########\n##########\n##########",
    "1 10\n1 10",
  ],
])("logic(%s) -> %s", (arg, expected) => {
  expect(logic(arg)).toBe(expected);
});
