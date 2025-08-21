// Package ops provides tools for composing expressions and math functions over quantities.
//
// There are two major units:
//
//   - Expression builder — a goal-directed way to combine quantities. You
//     declare the dimension you want (velocity, pressure, power, …) and then
//     build an expression from inputs until it resolves to that target. This
//     keeps calculations task-oriented and makes intent explicit.
//
//   - Math functions — familiar operations like exp, log, and the trig family,
//     but extended to understand units. They enforce physical rules (e.g. ln()
//     only accepts dimensionless ratios, sin() only accepts angles) so that
//     calculations stay meaningful.
//
// Together these make it possible to write calculations that look (reasonably)
// natural, read clearly, and remain dimensionally safe.
package ops
