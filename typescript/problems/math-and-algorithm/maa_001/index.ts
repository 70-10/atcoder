import { readFileSync } from "fs";
import { logic } from "./logic";

const arg = parseInt(readFileSync("/dev/stdin", "utf8").trim());

console.log(logic(arg));
