package GOlobalConfig

import (
	"strconv"
	"sync"
)

//---------------------------------------------------------------------------------
type Float struct {
	val float64
}

type Integer struct {
	val int64
}

type String struct {
	val string
}

type Bool struct {
	val bool
}

type face interface {
	toString() string
}

//---------------------------------------------------------------------------------
var mtx *sync.RWMutex = new(sync.RWMutex)
var config map[string]face = make(map[string]face)

//---------------------------------------------------------------------------------
func init() {

}

//---------------------------------------------------------------------------------
func AddString(key string, val string) bool {
	mtx.Lock()
	defer mtx.Unlock()

	if config[key] == nil {
		config[key] = &String{val}
		return true
	}
	return false
}

func GetString(key string) (string, bool) {
	mtx.RLock()
	defer mtx.RUnlock()

	if config[key] == nil {
		return "", false
	}

	return config[key].toString(), true
}

func (s *String) toString() string {
	return s.val
}

//---------------------------------------------------------------------------------
func AddInteger(key string, val int64) bool {
	mtx.Lock()
	defer mtx.Unlock()

	if config[key] == nil {
		config[key] = &Integer{val}
		return true
	}
	return false
}

func GetInteger(key string) (int64, bool) {
	mtx.RLock()
	defer mtx.RUnlock()

	if config[key] == nil {
		return 0, false
	}
	return config[key].(*Integer).val, true

}

func (i *Integer) toString() string {
	return strconv.FormatInt(i.val, 10)
}

//---------------------------------------------------------------------------------
func AddFloat(key string, val float64) bool {
	mtx.Lock()
	defer mtx.Unlock()

	if config[key] == nil {
		config[key] = &Float{val}
		return true
	}
	return false
}

func GetFloat(key string) (float64, bool) {
	mtx.RLock()
	defer mtx.RUnlock()

	if config[key] == nil {
		return 0.0, false
	}
	return config[key].(*Float).val, true
}

func (f *Float) toString() string {
	return strconv.FormatFloat(f.val, 'f', -1, 64)
}

//---------------------------------------------------------------------------------
func AddBool(key string, val bool) bool {
	mtx.Lock()
	defer mtx.Unlock()

	if config[key] == nil {
		config[key] = &Bool{val}
		return true
	}

	return false
}

func GetBool(key string) (bool, bool) {
	mtx.RLock()
	defer mtx.RUnlock()

	if config[key] == nil {
		return false, false
	}
	return config[key].(*Bool).val, true
}

func (b *Bool) toString() string {
	return strconv.FormatBool(b.val)
}
