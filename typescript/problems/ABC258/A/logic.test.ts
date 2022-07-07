import { logic } from "./logic";

test("63", () => {
  expect(logic(63)).toBe("22:03");
});
test("45", () => {
  expect(logic(45)).toBe("21:45");
});
