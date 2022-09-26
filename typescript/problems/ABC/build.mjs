import { execa } from "execa";

const problem_num = process.env.ABC;

execa("npx", [
  "esbuild",
  `problems/ABC/ABC${problem_num}/index.ts`,
  "--bundle",
  "--minify",
  "--platform=node",
  `--outfile=problems/ABC/ABC${problem_num}/index.js`,
]);
