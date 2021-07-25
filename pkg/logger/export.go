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

package logger

import (
	"fmt"
	"net/url"
	"ops/pkg/utils"
	"time"
)

// Debugf is the wrapper of default logger Debugf.
func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

// Infof is the wrapper of default logger Infof.
func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

// Warnf is the wrapper of default logger Warnf.
func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

// Errorf is the wrapper of default logger Errorf.
func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

// Debug is the wrapper of default logger Debug.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info is the wrapper of default logger Info.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn is the wrapper of default logger Warn.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error is the wrapper of default logger Error.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal is the wrapper of default logger Fatal.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf is the wrapper of default logger Fatalf.
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

// Sync syncs all logs, must be called after calling Init().
func Sync() {
	defaultLogger.Sync()
	stderrLogger.Sync()
	gressLogger.Sync()
	//httpFilterAccessLogger.Sync()
	//httpFilterDumpLogger.Sync()
	//restAPILogger.Sync()
}

//func DefaultLogger()*zap.SugaredLogger{
//	return defaultLogger
//}

// APIAccess logs admin api log.
func APIAccess(
	method, remoteAddr, path string,
	code int,
	bodyBytedReceived, bodyBytesSent int64,
	requestTime time.Time,
	processTime time.Duration) {
	entry := fmt.Sprintf("%s %s %s %v rx:%dB tx:%dB start:%v process:%v",
		method, remoteAddr, path, code,
		bodyBytedReceived, bodyBytesSent,
		requestTime.Format(time.RFC3339), processTime)

	//restAPILogger.Debug(entry)
	defaultLogger.Debug(entry)
}

// HTTPAccess logs http access log.
func HTTPAccess(line string) {
	httpFilterAccessLogger.Debug(line)
}

// NginxHTTPAccess is DEPRECATED, replaced by HTTPAccess.
func NginxHTTPAccess(remoteAddr, proto, method, path, referer, agent, realIP string,
	code int, bodyBytesSent int64,
	requestTime time.Duration, upstreamResponseTime time.Duration,
	upstreamAddr string, upstreamCode int, clientWriteBodyTime, clientReadBodyTime,
	routeTime time.Duration) {
	// mock nginx log_format:
	// '$remote_addr - $remote_user [$time_local] "$request" '
	// '$status $body_bytes_sent "$http_referer" '
	// '"$http_user_agent" "$http_x_forwarded_for" '
	// '$request_time $upstream_response_time $upstream_addr $upstream_status $pipe '
	// '$client_write_body_time' '$client_read_body_time' '$route_time';

	if referer == "" {
		referer = "-"
	}

	if agent == "" {
		agent = "-"
	} else {
		if a, err := url.QueryUnescape(agent); err == nil {
			agent = a
		}
	}

	if realIP == "" {
		realIP = "-"
	}

	if upstreamAddr == "" {
		upstreamAddr = "-"
	} else {
		if addr, err := url.QueryUnescape(upstreamAddr); err == nil {
			upstreamAddr = addr
		}
	}

	line := fmt.Sprintf(
		`%v - - [%v] "%s %s %s" `+
			`%v %v "%s" `+
			`"%s" "%s" `+
			`%f %f %v %v . `+
			`%f %f %f`,
		remoteAddr, utils.Now().Local(), method, path, proto,
		code, bodyBytesSent, referer,
		agent, realIP,
		requestTime.Seconds(), upstreamResponseTime.Seconds(), upstreamAddr, upstreamCode,
		clientWriteBodyTime.Seconds(), clientReadBodyTime.Seconds(), routeTime.Seconds())

	httpFilterAccessLogger.Debug(line)
}
