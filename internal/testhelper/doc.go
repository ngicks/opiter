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
package testhelper
