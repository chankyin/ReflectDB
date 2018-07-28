/*
 * reflectdb
 *
 * Copyright (C) 2018 chankyin
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package util

import (
	"errors"
	"os"
)

func CreateDirectory(dir string, perm os.FileMode) (alreadyExisted bool, err error) {
	stat, err := os.Stat(dir)
	if err != nil && !os.IsNotExist(err) {
		return
	}

	if err == nil {
		if stat.IsDir() {
			return true, nil
		} else {
			return false, errors.New(dir + " is a file")
		}
	}

	err = os.MkdirAll(dir, perm)
	return false, err
}
