import { readFileSync } from "fs";
import { logic } from "./logic";

const arg = readFileSync("/dev/stdin", "utf8").trim();

console.log(logic(arg));
