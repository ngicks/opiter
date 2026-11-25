// Package testhelper defines internal test helpers.
//
// The helpers assert the following iterator contracts:
//
//   - well-behaved: an iterator never calls yield again once yield returned
//     false.
//   - stateless: pure; every invocation of the same iterator value replays the
//     same data, unaffected by earlier breaks.
//   - stateful: single-use; data is consumed as it is yielded. A value
//     delivered to the yield call that rejected it (returned false) counts as
//     consumed: re-invoking the iterator resumes right after it, and a fully
//     consumed iterator yields nothing.
//   - single-use: only the first invocation is specified. The helpers never
//     re-invoke the same iterator value, because standard library iterators
//     documented as single-use disagree on what a second invocation yields.
package testhelper
