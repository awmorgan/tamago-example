// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "stringer -type NeighborState"; DO NOT EDIT.

package stack

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Incomplete-1]
	_ = x[Reachable-2]
	_ = x[Stale-3]
	_ = x[Delay-4]
	_ = x[Probe-5]
	_ = x[Static-6]
	_ = x[Unreachable-7]
}

const _NeighborState_name = "UnknownIncompleteReachableStaleDelayProbeStaticUnreachable"

var _NeighborState_index = [...]uint8{0, 7, 17, 26, 31, 36, 41, 47, 58}

func (i NeighborState) String() string {
	if i >= NeighborState(len(_NeighborState_index)-1) {
		return "NeighborState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NeighborState_name[_NeighborState_index[i]:_NeighborState_index[i+1]]
}
