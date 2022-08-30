import { execa } from "execa";

const problem_num = process.env.MAA;

execa("npx", [
  "esbuild",
  `problems/math-and-algorithm/maa_${problem_num}/index.ts`,
  "--bundle",
  "--minify",
  "--platform=node",
  `--outfile=problems/math-and-algorithm/maa_${problem_num}/index.js`,
]);
