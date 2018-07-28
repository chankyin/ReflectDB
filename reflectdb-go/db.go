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

package reflectdb

import (
	"github.com/chankyin/reflectdb/reflectdb-go/util"
	"github.com/theckman/go-flock"
	"path"
)

type Reflectdb struct {
	options    DBOptions
	schemaLock *flock.Flock
}

func NewReflectdb(options DBOptions) *Reflectdb {
	return &Reflectdb{
		options:    options,
		schemaLock: flock.NewFlock(path.Join(options.dir, "reflectdb.json")),
	}
}

// attempt to acquire a shared lock on the schema lock file
func (db *Reflectdb) Connect() (err error) {
	util.BlockingTimeout(db.schemaLock.Lock, db.schemaLock.Unlock, db.options.acquireTimeout)
	onComplete := make(chan error)
	onTimeout := make(chan bool)
	go func() {
		err := db.schemaLock.Lock()
		onComplete <- err
		select {
		case <-onTimeout:
		default:
			db.schemaLock.Unlock()
		}
	}()

	select {
	case err = <-onComplete:
		return err
	}
}
