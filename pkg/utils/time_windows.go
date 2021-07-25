/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
"syscall"
"time"
)

var(
	RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"
)

func Now() time.Time {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return time.Unix(0, tv.Nanoseconds())  //for windows
}

func NowUnixNano() int64 {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return tv.Nano()    //for windows
}

func NowUnix() int64 {
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	unixSecond, _ := tv.Unix()
	return unixSecond
}

func Since(t time.Time) time.Duration {
	return Now().Sub(t)
}

