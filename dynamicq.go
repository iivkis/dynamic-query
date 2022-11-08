package dynamicq

import (
	"strings"
)

type Dynamic struct {
	params []string
	args   []interface{}
}

func (d *Dynamic) Add(param string, v ...interface{}) {
	d.params = append(d.params, param)
	d.args = append(d.args, v...)
}

func (d *Dynamic) AddParam(param string) {
	d.params = append(d.params, param)
}

func (d *Dynamic) AddArg(v ...interface{}) {
	d.args = append(d.args, v...)
}

func (d *Dynamic) Glue(query *string) {
	p := strings.Join(d.params, " AND ")
	if len(p) > 0 {
		p = " WHERE " + p
	}
	*query += p
}

func (d *Dynamic) Attr(query *string, attr string) {
	*query += " " + attr
}

func (d *Dynamic) Limit(query *string, limit int64) {
	if limit != 0 {
		*query += " LIMIT ?"
		d.AddArg(limit)
	}
}

func (d *Dynamic) Offset(query *string, offset int64) {
	if offset != 0 {
		*query += " OFFSET ?"
		d.AddArg(offset)
	}
}

func (d *Dynamic) Args() []interface{} {
	return d.args
}
