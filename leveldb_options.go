package ezdb

import "github.com/syndtr/goleveldb/leveldb/opt"

type LevelDBOptions struct {
	Open  *opt.Options
	Read  *opt.ReadOptions
	Write *opt.WriteOptions
}

func (o *LevelDBOptions) GetOpen() *opt.Options {
	if o == nil {
		return nil
	}
	return o.Open
}

func (o *LevelDBOptions) GetRead() *opt.ReadOptions {
	if o == nil {
		return nil
	}
	return o.Read
}

func (o *LevelDBOptions) GetWrite() *opt.WriteOptions {
	if o == nil {
		return nil
	}
	return o.Write
}
