//go:build linux && musl

/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package metadata

// fetchGlibcVersion returns an empty string on musl-based systems (e.g. OpenWrt,
// Alpine).
//
// The upstream implementation calls gnu_get_libc_version(3) via cgo, which needs
// <gnu/libc-version.h> -- a glibc-only header that musl does not provide, so the
// cgo variant cannot be compiled against a musl toolchain at all.
//
// There is no musl equivalent worth reporting here: musl exposes no runtime
// version query, and the field is named after glibc specifically. Reporting
// nothing is more honest than inventing a value.
func (c *fetchConfig) fetchGlibcVersion() string {
	return ""
}
