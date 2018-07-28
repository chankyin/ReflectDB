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
	"time"
)

func BlockingTimeout(executor func() error, undo func() error, timeout time.Duration) (executed bool, err error) {
	onComplete := make(chan error)
	onTimeout := make(chan bool, 1)
	go func() {
		err := executor()
		select {
		case <-onTimeout:
			if err != nil {
				panic(err) // nobody to receive the error
			}
			err = undo()
			if err != nil {
				panic(err) // nobody to receive the error
			}
		default:
			onComplete <- err
		}
	}()
	select {
	case err = <-onComplete:
		return true, err
	case <-time.After(timeout):
		close(onTimeout)
		return false, nil
	}
}
