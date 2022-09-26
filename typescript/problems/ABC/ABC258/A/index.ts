import { readFileSync } from "fs";
import { logic } from "./logic";

const K = parseInt(readFileSync("/dev/stdin", "utf8").trim());

console.log(logic(K));
